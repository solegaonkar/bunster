package generator

import (
	"fmt"

	"github.com/yassinebenaid/bunster/ast"
	"github.com/yassinebenaid/bunster/ir"
)

func (g *generator) handleArithmeticCommand(buf *InstructionBuffer, cmd ast.ArithmeticCommand) {
	var cmdbuf InstructionBuffer

	cmdbuf.add(ir.CloneStreamManager{})
	g.handleRedirections(&cmdbuf, cmd.Redirections)

	cmdbuf.add(ir.Declare{Name: "arithmeticResult", Value: ir.Literal("0")})

	for _, arithmetic := range cmd.Arithmetic {
		cmdbuf.add(ir.Set{Name: "arithmeticResult", Value: g.handleArithmeticExpression(&cmdbuf, arithmetic)})
	}

	cmdbuf.add(ir.Literal("if arithmeticResult == 0 { shell.ExitCode = 1 } else { shell.ExitCode = 0  }\n"))

	*buf = append(*buf, ir.Closure(cmdbuf))
}

func (g *generator) handleArithmeticSubstitution(buf *InstructionBuffer, expr ast.Arithmetic) ir.Instruction {
	var cmdbuf InstructionBuffer

	cmdbuf.add(ir.Declare{Name: "arithmeticResult", Value: ir.Literal("0")})

	for _, arithmetic := range expr {
		cmdbuf.add(ir.Set{Name: "arithmeticResult", Value: g.handleArithmeticExpression(&cmdbuf, arithmetic)})
	}

	cmdbuf.add(ir.Literal("return runtime.FormatInt(arithmeticResult), shell.ExitCode"))

	name := fmt.Sprintf("expr%d", g.expressionsCount)
	buf.add(ir.ExpressionClosure{
		Body: cmdbuf,
		Name: name,
	})

	return ir.Literal(name)
}

func (g *generator) handleArithmeticExpression(buf *InstructionBuffer, expr ast.Expression) ir.Instruction {
	switch v := expr.(type) {
	case ast.PostIncDecArithmetic:
		return (ir.VarIncDec{Operand: v.Operand, Operator: v.Operator, Post: true})
	case ast.PreIncDecArithmetic:
		return (ir.VarIncDec{Operand: v.Operand, Operator: v.Operator})
	case ast.Unary:
		return (ir.UnaryArithmetic{Operand: g.handleArithmeticExpression(buf, v.Operand), Operator: v.Operator})
	case ast.BitFlip:
		return (ir.UnaryArithmetic{Operand: g.handleArithmeticExpression(buf, v.Operand), Operator: "^"})
	case ast.Negation:
		return (ir.NegateArithmetic{Value: g.handleArithmeticExpression(buf, v.Operand)})
	case ast.Binary:
		switch v.Operator {
		case "+", "-", "*", "/", "%", "|", "&", "^", "<<", ">>":
			return (ir.BinaryArithmetic{
				Left:     g.handleArithmeticExpression(buf, v.Left),
				Right:    g.handleArithmeticExpression(buf, v.Right),
				Operator: v.Operator,
			})
		case "**":
			return (ir.IntPower{
				Operand: g.handleArithmeticExpression(buf, v.Left),
				Pow:     g.handleArithmeticExpression(buf, v.Right),
			})
		case ">", "<", ">=", "<=", "!=", "==", "&&", "||":
			return (ir.CompareInt{
				Left:     g.handleArithmeticExpression(buf, v.Left),
				Right:    g.handleArithmeticExpression(buf, v.Right),
				Operator: v.Operator,
			})
		case "=":
			name := string(v.Left.(ast.Var))
			buf.add(ir.SetVar{
				Key:   name,
				Value: ir.FormatInt{Value: g.handleArithmeticExpression(buf, v.Right)},
			})

			return ir.ParseInt{Value: ir.ReadVar(name)}
		case "+=", "-=", "*=", "/=", "%=", "|=", "&=", "^=", "<<=", ">>=":
			name := string(v.Left.(ast.Var))
			buf.add(ir.SetVar{
				Key: name,
				Value: ir.FormatInt{Value: ir.BinaryArithmetic{
					Left:     g.handleArithmeticExpression(buf, v.Left),
					Right:    g.handleArithmeticExpression(buf, v.Right),
					Operator: v.Operator[:len(v.Operator)-1],
				}},
			})

			return ir.ParseInt{Value: ir.ReadVar(name)}
		default:
			panic("unsupported binary arithmetic: " + v.Operator)
		}
	case ast.Conditional:
		return (ir.ConditionalInt{
			Test:      g.handleArithmeticExpression(buf, v.Test),
			Body:      g.handleArithmeticExpression(buf, v.Body),
			Alternate: g.handleArithmeticExpression(buf, v.Alternate),
		})
	case ast.Number:
		return ir.Literal(v)
	case ast.Var:
		return ir.ParseInt{Value: ir.ReadVar(v)}
	default:
		panic(fmt.Sprintf("unsupported arithmetic expression: %T", v))
	}
}
