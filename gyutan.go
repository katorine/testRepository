//
package main

//
import "fmt"

const value_verylikely = 5   // 表情種類の値5 とても強い
const value_likely = 4       // 表情種類の値4 強い
const value_possible = 3     // 表情種類の値3 普通
const value_unlikely = 2     // 表情種類の値2 弱い
const value_veryunlikely = 1 // 表情種類の値1 とても弱い

const type_joy = 1           // 表情の種類 たのしい
const type_sorrow = 2        // 表情の種類 かなしい
const type_anger = 3         // 表情の種類 いかり
const type_suprise = 4       // 表情の種類 おどろき
const type_loop_end = 5      // 表情の種類 終了

// ★TODO
const url_joy = todo         // URL値 たのしい
const url_sorrow = todo      // URL値 かなしい
const url_anger = todo       // URL値 いかり
const url_suprise = todo     // URL値 おどろき
const url_unknown = todo     // URL値 不明

// 表情種類の値を返す
func transType( type string ) int{
  switch (type) {
    case VERY_LIKELY :
      return value_verylikely
    case LIKELY :
      return value_likely
    case POSSIBLE :
      return value_possible
    case UNLIKELY :
      return value_unlikely
    default :
      return value_veryunlikely
  }
}

//関数はfuncキーワードで定義
func main(){

  // 表情種類最大値
  var maxValueOfType int := value_veryunlikely
  // 表情種類最大値重複判定
  var chofukuCheck bool
  // 選択された表情（URL）
  var selectedTypeURL string := url_joy
  // 表情種類のmap
  typeMap := map[int]int{}
  // 表情種類のURLmap
  typeUrlMap := map[int]string{}

// ★TODO
  // 表情種類のmap設定
  typeMap[type_joy] = transType( getJson )           // TODO:getJson joyの値
  typeMap[type_sorrow] = transType( getJson )        // TODO:getJson sorrowの値
  typeMap[type_anger] = transType( getJson )         // TODO:getJson angerの値
  typeMap[type_suprise] = transType( getJson )       // TODO:getJson supriseの値

  // 表情種類のURLmap設定
  typeUrlMap[type_joy] = url_joy
  typeUrlMap[type_sorrow] = url_sorrow
  typeUrlMap[type_anger] = url_anger
  typeUrlMap[type_suprise] = url_suprise
  
  for typeKind := type_joy ; typeKind < type_loop_end ; typeKind++ {
    // 表情最大値比較からの表情選択処理
    if maxValueOfType == typeMap[typeKind] {
      // 表情種類最大値重複判定 （重複あり）
      chofukuCheck = true
    } else if maxValueOfType < typeMap[typeKind] {
      // 表情種類最大値変更
      maxValueOfType = typeMap[typeKind]
      // 選択された表情変更
      selectedTypeURL = typeUrlMap[typeKind]
      // 表情種類最大値重複判定 （重複設定初期化）
      chofukuCheck = false
    } 
  }

  // 重複ありの場合、不明を設定
  if chofukuCheck {
    selectedTypeURL = url_unknown
  } 

// ★TODO
  // 
  return           // TODO:
}