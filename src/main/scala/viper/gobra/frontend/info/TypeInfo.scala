// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2011-2020 ETH Zurich.

package viper.gobra.frontend.info

import org.bitbucket.inkytonik.kiama.relation.Tree
import viper.gobra.ast.frontend._
import viper.gobra.frontend.info.base.SymbolTable.Regular
import viper.gobra.frontend.info.base.Type.Type
import viper.gobra.theory.Addressability

trait TypeInfo extends ExternalTypeInfo {

  def context: Info.Context

  def typOfExprOrType(expr: PExpressionOrType): Type
  def addressability(expr: PExpression): Addressability
  def addressableVar(id: PIdnNode): Addressability

  def codeRoot(n: PNode): PScope

  def tree: Tree[PNode, PPackage]

  def regular(n: PIdnNode): Regular

  def isDef(n: PIdnUnk): Boolean

  def resolve(n: PExpressionOrType): Option[AstPattern.Pattern]
  def exprOrType(n: PExpressionOrType): Either[PExpression, PType]

  def nilType(n: PNilLit): Option[Type]
}
