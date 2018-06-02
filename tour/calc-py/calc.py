import sys
#from antlr4 import *
from antlr4 import CommonTokenStream, FileStream
from antlr4.InputStream import InputStream
from LabeledExpr.LabeledExprLexer import LabeledExprLexer
from LabeledExpr.LabeledExprParser import LabeledExprParser
from LabeledExpr.LabeledExprVisitor import LabeledExprVisitor


class EvalVisitor(LabeledExprVisitor):
    def __init__(self):
        self.memory = {}

    def visitAssign(self, ctx):
        name = ctx.ID().getText()
        value = self.visit(ctx.expr())
        self.memory[name] = value
        return value

    def visitPrintExpr(self, ctx):
        value = self.visit(ctx.expr())
        print(value)
        return 0

    def visitInt(self, ctx):
        return ctx.INT().getText()

    def visitId(self, ctx):
        name = ctx.ID().getText()
        if name in self.memory:
            return self.memory[name]
        return 0

    def visitMulDiv(self, ctx):
        left = int(self.visit(ctx.expr(0)))
        right = int(self.visit(ctx.expr(1)))
        if ctx.op.type == LabeledExprParser.MUL:
            return left * right
        return left / right

    def visitAddSub(self, ctx):
        left = int(self.visit(ctx.expr(0)))
        right = int(self.visit(ctx.expr(1)))
        if ctx.op.type == LabeledExprParser.ADD:
            return left + right
        return left - right

    def visitParens(self, ctx):
        return self.visit(ctx.expr())


if __name__ == '__main__':
    if len(sys.argv) > 1:
        input_stream = FileStream(sys.argv[1])
    else:
        input_stream = InputStream(sys.stdin.readline())

    lexer = LabeledExprLexer(input_stream)
    token_stream = CommonTokenStream(lexer)
    parser = LabeledExprParser(token_stream)
    tree = parser.prog()

    # lisp_tree_str = tree.toStringTree(recog=parser)
    # print(lisp_tree_str)

    visitor = EvalVisitor()
    visitor.visit(tree)