// Copyright 2017 The Bazel Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package syntax provides a Skylark parser and abstract syntax tree.
package syntax

// A Node is a node in a Skylark syntax tree.
type Node interface {
	// Span returns the start and end position of the expression.
	Span() (start, end Position)

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	// Comments returns the comments associated with this node.
	// It returns nil if RetainComments was not specified during parsing,
	// or if AllocComments was not called.
	Comments() *Comments

	// AllocComments allocates a new Comments node if there was none.
	// This makes possible to add new comments using Comments() method.
	AllocComments()
=======
	Comment() *Comments
>>>>>>> Attach comments to AST nodes.
=======
	AllocComments()
	Comments() *Comments
>>>>>>> Add CommentsRef to allow allocating comments
=======
	// Comments returns the comments associated with this node.
	// It returns nil if RetainComments was not specified during parsing,
	// or if AllocComments was not called.
	Comments() *Comments

	// AllocComments allocates a new Comments node if there was none.
	// This makes possible to add new comments using Comments() method.
	AllocComments()
>>>>>>> Remove the COMMENT tokens
}

// A Comment represents a single # comment.
type Comment struct {
<<<<<<< HEAD
<<<<<<< HEAD
	Start Position
	Text  string // without trailing newline
=======
	Start  Position
<<<<<<< HEAD
	Token  string // without trailing newline
<<<<<<< HEAD
	Suffix bool   // an end of line (not whole line) comment
>>>>>>> Attach comments to AST nodes.
=======
=======
	Text   string // without trailing newline
>>>>>>> Remove global constant keepComments
	Suffix bool   // an end-of-line (not whole line) comment
>>>>>>> - Rename flattenAST
=======
	Start Position
	Text  string // without trailing newline
>>>>>>> - Removed the .Suffix boolean
}

// Comments collects the comments associated with an expression.
type Comments struct {
	Before []Comment // whole-line comments before this expression
<<<<<<< HEAD
	Suffix []Comment // end-of-line comments after this expression (up to 1)
=======
	Suffix []Comment // end-of-line comments after this expression
>>>>>>> Attach comments to AST nodes.

	// For top-level expressions only, After lists whole-line
	// comments following the expression.
	After []Comment
}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
// A commentsRef is a possibly-nil reference to a set of comments.
// A commentsRef is embedded in each type of syntax node,
// and provides its Comments and AllocComments methods.
type commentsRef struct{ ref *Comments }

// Comments returns the comments associated with a syntax node,
// or nil if AllocComments has not yet been called.
func (cr commentsRef) Comments() *Comments { return cr.ref }

// AllocComments enables comments to be associated with a syntax node.
func (cr *commentsRef) AllocComments() {
	if cr.ref == nil {
		cr.ref = new(Comments)
	}
=======
// Comment returns the receiver. This isn't useful by itself, but
// a Comments struct is embedded into all the expression
// implementation types, and this gives each of those a Comment
// method to satisfy the Expr interface.
func (c *Comments) Comment() *Comments {
	return c
>>>>>>> Attach comments to AST nodes.
=======
// A CommentsRef is a possibly-nil reference to a set of comments.
// A CommentsRef is embedded in each type of syntax node,
=======
// A commentsRef is a possibly-nil reference to a set of comments.
// A commentsRef is embedded in each type of syntax node,
>>>>>>> - Removed the .Suffix boolean
// and provides its Comments and AllocComments methods.
type commentsRef struct{ ref *Comments }

// Comments returns the comments associated with a syntax node,
// or nil if AllocComments has not yet been called.
func (cr commentsRef) Comments() *Comments { return cr.ref }

// AllocComments enables comments to be associated with a syntax node.
func (cr *commentsRef) AllocComments() {
	if cr.ref == nil {
		cr.ref = new(Comments)
	}
>>>>>>> Add CommentsRef to allow allocating comments
}

// Start returns the start position of the expression.
func Start(n Node) Position {
	start, _ := n.Span()
	return start
}

// End returns the end position of the expression.
func End(n Node) Position {
	_, end := n.Span()
	return end
}

// A File represents a Skylark file.
type File struct {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	commentsRef
=======
	Comments
>>>>>>> Attach comments to AST nodes.
=======
	*Comments
>>>>>>> - Rename flattenAST
=======
	CommentsRef
>>>>>>> Add CommentsRef to allow allocating comments
=======
	commentsRef
>>>>>>> - Removed the .Suffix boolean
	Path  string
	Stmts []Stmt

	// set by resolver:
	Locals []*Ident // this file's (comprehension-)local variables
}

func (x *File) Span() (start, end Position) {
	if len(x.Stmts) == 0 {
		return
	}
	start, _ = x.Stmts[0].Span()
	_, end = x.Stmts[len(x.Stmts)-1].Span()
	return start, end
}

// A Stmt is a Skylark statement.
type Stmt interface {
	Node
	stmt()
}

func (*AssignStmt) stmt() {}
func (*BranchStmt) stmt() {}
func (*DefStmt) stmt()    {}
func (*ExprStmt) stmt()   {}
func (*ForStmt) stmt()    {}
func (*IfStmt) stmt()     {}
func (*LoadStmt) stmt()   {}
func (*ReturnStmt) stmt() {}

// An AssignStmt represents an assignment:
//	x = 0
//	x, y = y, x
// 	x += 1
type AssignStmt struct {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	commentsRef
=======
	Comments
>>>>>>> Attach comments to AST nodes.
=======
	*Comments
>>>>>>> - Rename flattenAST
=======
	CommentsRef
>>>>>>> Add CommentsRef to allow allocating comments
=======
	commentsRef
>>>>>>> - Removed the .Suffix boolean
	OpPos Position
	Op    Token // = EQ | {PLUS,MINUS,STAR,PERCENT}_EQ
	LHS   Expr
	RHS   Expr
}

func (x *AssignStmt) Span() (start, end Position) {
	start, _ = x.LHS.Span()
	_, end = x.RHS.Span()
	return
}

// A Function represents the common parts of LambdaExpr and DefStmt.
type Function struct {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	commentsRef
=======
	Comments
>>>>>>> Attach comments to AST nodes.
=======
	*Comments
>>>>>>> - Rename flattenAST
=======
	CommentsRef
>>>>>>> Add CommentsRef to allow allocating comments
=======
	commentsRef
>>>>>>> - Removed the .Suffix boolean
	StartPos Position // position of DEF or LAMBDA token
	Params   []Expr   // param = ident | ident=expr | *ident | **ident
	Body     []Stmt

	// set by resolver:
	HasVarargs bool     // whether params includes *args (convenience)
	HasKwargs  bool     // whether params includes **kwargs (convenience)
	Locals     []*Ident // this function's local variables, parameters first
	FreeVars   []*Ident // enclosing local variables to capture in closure
}

func (x *Function) Span() (start, end Position) {
	_, end = x.Body[len(x.Body)-1].Span()
	return x.StartPos, end
}

// A DefStmt represents a function definition.
type DefStmt struct {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	commentsRef
=======
	Comments
>>>>>>> Attach comments to AST nodes.
=======
	*Comments
>>>>>>> - Rename flattenAST
=======
	CommentsRef
>>>>>>> Add CommentsRef to allow allocating comments
=======
	commentsRef
>>>>>>> - Removed the .Suffix boolean
	Def  Position
	Name *Ident
	Function
}

func (x *DefStmt) Span() (start, end Position) {
	_, end = x.Function.Body[len(x.Body)-1].Span()
	return x.Def, end
}

// An ExprStmt is an expression evaluated for side effects.
type ExprStmt struct {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	commentsRef
=======
	Comments
>>>>>>> Attach comments to AST nodes.
=======
	*Comments
>>>>>>> - Rename flattenAST
=======
	CommentsRef
>>>>>>> Add CommentsRef to allow allocating comments
=======
	commentsRef
>>>>>>> - Removed the .Suffix boolean
	X Expr
}

func (x *ExprStmt) Span() (start, end Position) {
	return x.X.Span()
}

// An IfStmt is a conditional: If Cond: True; else: False.
// 'elseif' is desugared into a chain of IfStmts.
type IfStmt struct {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	commentsRef
=======
	Comments
>>>>>>> Attach comments to AST nodes.
=======
	*Comments
>>>>>>> - Rename flattenAST
=======
	CommentsRef
>>>>>>> Add CommentsRef to allow allocating comments
=======
	commentsRef
>>>>>>> - Removed the .Suffix boolean
	If      Position // IF or ELIF
	Cond    Expr
	True    []Stmt
	ElsePos Position // ELSE or ELIF
	False   []Stmt   // optional
}

func (x *IfStmt) Span() (start, end Position) {
	body := x.False
	if body == nil {
		body = x.True
	}
	_, end = body[len(body)-1].Span()
	return x.If, end
}

// A LoadStmt loads another module and binds names from it:
// load(Module, "x", y="foo").
//
// The AST is slightly unfaithful to the concrete syntax here because
// Skylark's load statement, so that it can be implemented in Python,
// binds some names (like y above) with an identifier and some (like x)
// without.  For consistency we create fake identifiers for all the
// strings.
type LoadStmt struct {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	commentsRef
=======
	Comments
>>>>>>> Attach comments to AST nodes.
=======
	*Comments
>>>>>>> - Rename flattenAST
=======
	CommentsRef
>>>>>>> Add CommentsRef to allow allocating comments
=======
	commentsRef
>>>>>>> - Removed the .Suffix boolean
	Load   Position
	Module *Literal // a string
	From   []*Ident // name defined in loading module
	To     []*Ident // name in loaded module
	Rparen Position
}

func (x *LoadStmt) Span() (start, end Position) {
	return x.Load, x.Rparen
}

// A BranchStmt changes the flow of control: break, continue, pass.
type BranchStmt struct {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	commentsRef
=======
	Comments
>>>>>>> Attach comments to AST nodes.
=======
	*Comments
>>>>>>> - Rename flattenAST
=======
	CommentsRef
>>>>>>> Add CommentsRef to allow allocating comments
=======
	commentsRef
>>>>>>> - Removed the .Suffix boolean
	Token    Token // = BREAK | CONTINUE | PASS
	TokenPos Position
}

func (x *BranchStmt) Span() (start, end Position) {
	return x.TokenPos, x.TokenPos.add(x.Token.String())
}

// A ReturnStmt returns from a function.
type ReturnStmt struct {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	commentsRef
=======
	Comments
>>>>>>> Attach comments to AST nodes.
=======
	*Comments
>>>>>>> - Rename flattenAST
=======
	CommentsRef
>>>>>>> Add CommentsRef to allow allocating comments
=======
	commentsRef
>>>>>>> - Removed the .Suffix boolean
	Return Position
	Result Expr // may be nil
}

func (x *ReturnStmt) Span() (start, end Position) {
	if x.Result == nil {
		return x.Return, x.Return.add("return")
	}
	_, end = x.Result.Span()
	return x.Return, end
}

// An Expr is a Skylark expression.
type Expr interface {
	Node
	expr()
}

func (*BinaryExpr) expr()    {}
func (*CallExpr) expr()      {}
func (*Comprehension) expr() {}
func (*CondExpr) expr()      {}
func (*DictEntry) expr()     {}
func (*DictExpr) expr()      {}
func (*DotExpr) expr()       {}
func (*Ident) expr()         {}
func (*IndexExpr) expr()     {}
func (*LambdaExpr) expr()    {}
func (*ListExpr) expr()      {}
func (*Literal) expr()       {}
func (*SliceExpr) expr()     {}
func (*TupleExpr) expr()     {}
func (*UnaryExpr) expr()     {}

// An Ident represents an identifier.
type Ident struct {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	commentsRef
=======
	Comments
>>>>>>> Attach comments to AST nodes.
=======
	*Comments
>>>>>>> - Rename flattenAST
=======
	CommentsRef
>>>>>>> Add CommentsRef to allow allocating comments
=======
	commentsRef
>>>>>>> - Removed the .Suffix boolean
	NamePos Position
	Name    string

	// set by resolver:

	Scope uint8 // one of resolve.{Undefined,Local,Free,Global,Builtin}
	Index int   // index into enclosing {DefStmt,File}.Locals (if scope==Local) or DefStmt.FreeVars (if scope==Free)
}

func (x *Ident) Span() (start, end Position) {
	return x.NamePos, x.NamePos.add(x.Name)
}

// A Literal represents a literal string or number.
type Literal struct {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	commentsRef
=======
	Comments
>>>>>>> Attach comments to AST nodes.
=======
	*Comments
>>>>>>> - Rename flattenAST
=======
	CommentsRef
>>>>>>> Add CommentsRef to allow allocating comments
=======
	commentsRef
>>>>>>> - Removed the .Suffix boolean
	Token    Token // = STRING | INT
	TokenPos Position
	Raw      string      // uninterpreted text
	Value    interface{} // = string | int64 | *big.Int
}

func (x *Literal) Span() (start, end Position) {
	return x.TokenPos, x.TokenPos.add(x.Raw)
}

// A CallExpr represents a function call expression: Fn(Args).
type CallExpr struct {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	commentsRef
=======
	Comments
>>>>>>> Attach comments to AST nodes.
=======
	*Comments
>>>>>>> - Rename flattenAST
=======
	CommentsRef
>>>>>>> Add CommentsRef to allow allocating comments
=======
	commentsRef
>>>>>>> - Removed the .Suffix boolean
	Fn     Expr
	Lparen Position
	Args   []Expr
	Rparen Position
}

func (x *CallExpr) Span() (start, end Position) {
	start, _ = x.Fn.Span()
	return start, x.Rparen.add(")")
}

// A DotExpr represents a field or method selector: X.Name.
type DotExpr struct {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	commentsRef
=======
	Comments
>>>>>>> Attach comments to AST nodes.
=======
	*Comments
>>>>>>> - Rename flattenAST
=======
	CommentsRef
>>>>>>> Add CommentsRef to allow allocating comments
=======
	commentsRef
>>>>>>> - Removed the .Suffix boolean
	X       Expr
	Dot     Position
	NamePos Position
	Name    *Ident
}

func (x *DotExpr) Span() (start, end Position) {
	start, _ = x.X.Span()
	_, end = x.Name.Span()
	return
}

// A Comprehension represents a list or dict comprehension:
// [Body for ... if ...] or {Body for ... if ...}
type Comprehension struct {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	commentsRef
=======
	Comments
>>>>>>> Attach comments to AST nodes.
=======
	*Comments
>>>>>>> - Rename flattenAST
=======
	CommentsRef
>>>>>>> Add CommentsRef to allow allocating comments
=======
	commentsRef
>>>>>>> - Removed the .Suffix boolean
	Curly   bool // {x:y for ...} or {x for ...}, not [x for ...]
	Lbrack  Position
	Body    Expr
	Clauses []Node // = *ForClause | *IfClause
	Rbrack  Position
}

func (x *Comprehension) Span() (start, end Position) {
	return x.Lbrack, x.Rbrack.add("]")
}

// A ForStmt represents a loop: for Vars in X: Body.
type ForStmt struct {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	commentsRef
=======
	Comments
>>>>>>> Attach comments to AST nodes.
=======
	*Comments
>>>>>>> - Rename flattenAST
=======
	CommentsRef
>>>>>>> Add CommentsRef to allow allocating comments
=======
	commentsRef
>>>>>>> - Removed the .Suffix boolean
	For  Position
	Vars Expr // name, or tuple of names
	X    Expr
	Body []Stmt
}

func (x *ForStmt) Span() (start, end Position) {
	_, end = x.Body[len(x.Body)-1].Span()
	return x.For, end
}

// A ForClause represents a for clause in a list comprehension: for Vars in X.
type ForClause struct {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	commentsRef
=======
	Comments
>>>>>>> Attach comments to AST nodes.
=======
	*Comments
>>>>>>> - Rename flattenAST
=======
	CommentsRef
>>>>>>> Add CommentsRef to allow allocating comments
=======
	commentsRef
>>>>>>> - Removed the .Suffix boolean
	For  Position
	Vars Expr // name, or tuple of names
	In   Position
	X    Expr
}

func (x *ForClause) Span() (start, end Position) {
	_, end = x.X.Span()
	return x.For, end
}

// An IfClause represents an if clause in a list comprehension: if Cond.
type IfClause struct {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	commentsRef
=======
	Comments
>>>>>>> Attach comments to AST nodes.
=======
	*Comments
>>>>>>> - Rename flattenAST
=======
	CommentsRef
>>>>>>> Add CommentsRef to allow allocating comments
=======
	commentsRef
>>>>>>> - Removed the .Suffix boolean
	If   Position
	Cond Expr
}

func (x *IfClause) Span() (start, end Position) {
	_, end = x.Cond.Span()
	return x.If, end
}

// A DictExpr represents a dictionary literal: { List }.
type DictExpr struct {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	commentsRef
=======
	Comments
>>>>>>> Attach comments to AST nodes.
=======
	*Comments
>>>>>>> - Rename flattenAST
=======
	CommentsRef
>>>>>>> Add CommentsRef to allow allocating comments
=======
	commentsRef
>>>>>>> - Removed the .Suffix boolean
	Lbrace Position
	List   []Expr // all *DictEntrys
	Rbrace Position
}

func (x *DictExpr) Span() (start, end Position) {
	return x.Lbrace, x.Rbrace.add("}")
}

// A DictEntry represents a dictionary entry: Key: Value.
// Used only within a DictExpr.
type DictEntry struct {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	commentsRef
=======
	Comments
>>>>>>> Attach comments to AST nodes.
=======
	*Comments
>>>>>>> - Rename flattenAST
=======
	CommentsRef
>>>>>>> Add CommentsRef to allow allocating comments
=======
	commentsRef
>>>>>>> - Removed the .Suffix boolean
	Key   Expr
	Colon Position
	Value Expr
}

func (x *DictEntry) Span() (start, end Position) {
	start, _ = x.Key.Span()
	_, end = x.Value.Span()
	return start, end
}

// A LambdaExpr represents an inline function abstraction.
//
// Although they may be added in future, lambda expressions are not
// currently part of the Skylark spec, so their use is controlled by the
// resolver.AllowLambda flag.
type LambdaExpr struct {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	commentsRef
=======
	Comments
>>>>>>> Attach comments to AST nodes.
=======
	*Comments
>>>>>>> - Rename flattenAST
=======
	CommentsRef
>>>>>>> Add CommentsRef to allow allocating comments
=======
	commentsRef
>>>>>>> - Removed the .Suffix boolean
	Lambda Position
	Function
}

func (x *LambdaExpr) Span() (start, end Position) {
	_, end = x.Function.Body[len(x.Body)-1].Span()
	return x.Lambda, end
}

// A ListExpr represents a list literal: [ List ].
type ListExpr struct {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	commentsRef
=======
	Comments
>>>>>>> Attach comments to AST nodes.
=======
	*Comments
>>>>>>> - Rename flattenAST
=======
	CommentsRef
>>>>>>> Add CommentsRef to allow allocating comments
=======
	commentsRef
>>>>>>> - Removed the .Suffix boolean
	Lbrack Position
	List   []Expr
	Rbrack Position
}

func (x *ListExpr) Span() (start, end Position) {
	return x.Lbrack, x.Rbrack.add("]")
}

// CondExpr represents the conditional: X if COND else ELSE.
type CondExpr struct {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	commentsRef
=======
	Comments
>>>>>>> Attach comments to AST nodes.
=======
	*Comments
>>>>>>> - Rename flattenAST
=======
	CommentsRef
>>>>>>> Add CommentsRef to allow allocating comments
=======
	commentsRef
>>>>>>> - Removed the .Suffix boolean
	If      Position
	Cond    Expr
	True    Expr
	ElsePos Position
	False   Expr
}

func (x *CondExpr) Span() (start, end Position) {
	start, _ = x.True.Span()
	_, end = x.False.Span()
	return start, end
}

// A TupleExpr represents a tuple literal: (List).
type TupleExpr struct {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	commentsRef
=======
	Comments
>>>>>>> Attach comments to AST nodes.
=======
	*Comments
>>>>>>> - Rename flattenAST
=======
	CommentsRef
>>>>>>> Add CommentsRef to allow allocating comments
=======
	commentsRef
>>>>>>> - Removed the .Suffix boolean
	Lparen Position // optional (e.g. in x, y = 0, 1), but required if List is empty
	List   []Expr
	Rparen Position
}

func (x *TupleExpr) Span() (start, end Position) {
	if x.Lparen.IsValid() {
		return x.Lparen, x.Rparen
	} else {
		return Start(x.List[0]), End(x.List[len(x.List)-1])
	}
}

// A UnaryExpr represents a unary expression: Op X.
type UnaryExpr struct {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	commentsRef
=======
	Comments
>>>>>>> Attach comments to AST nodes.
=======
	*Comments
>>>>>>> - Rename flattenAST
=======
	CommentsRef
>>>>>>> Add CommentsRef to allow allocating comments
=======
	commentsRef
>>>>>>> - Removed the .Suffix boolean
	OpPos Position
	Op    Token
	X     Expr
}

func (x *UnaryExpr) Span() (start, end Position) {
	_, end = x.X.Span()
	return x.OpPos, end
}

// A BinaryExpr represents a binary expression: X Op Y.
type BinaryExpr struct {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	commentsRef
=======
	Comments
>>>>>>> Attach comments to AST nodes.
=======
	*Comments
>>>>>>> - Rename flattenAST
=======
	CommentsRef
>>>>>>> Add CommentsRef to allow allocating comments
=======
	commentsRef
>>>>>>> - Removed the .Suffix boolean
	X     Expr
	OpPos Position
	Op    Token
	Y     Expr
}

func (x *BinaryExpr) Span() (start, end Position) {
	start, _ = x.X.Span()
	_, end = x.Y.Span()
	return start, end
}

// A SliceExpr represents a slice or substring expression: X[Lo:Hi:Step].
type SliceExpr struct {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	commentsRef
=======
	Comments
>>>>>>> Attach comments to AST nodes.
=======
	*Comments
>>>>>>> - Rename flattenAST
=======
	CommentsRef
>>>>>>> Add CommentsRef to allow allocating comments
=======
	commentsRef
>>>>>>> - Removed the .Suffix boolean
	X            Expr
	Lbrack       Position
	Lo, Hi, Step Expr // all optional
	Rbrack       Position
}

func (x *SliceExpr) Span() (start, end Position) {
	start, _ = x.X.Span()
	return start, x.Rbrack
}

// An IndexExpr represents an index expression: X[Y].
type IndexExpr struct {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	commentsRef
=======
	Comments
>>>>>>> Attach comments to AST nodes.
=======
	*Comments
>>>>>>> - Rename flattenAST
=======
	CommentsRef
>>>>>>> Add CommentsRef to allow allocating comments
=======
	commentsRef
>>>>>>> - Removed the .Suffix boolean
	X      Expr
	Lbrack Position
	Y      Expr
	Rbrack Position
}

func (x *IndexExpr) Span() (start, end Position) {
	start, _ = x.X.Span()
	return start, x.Rbrack
}
