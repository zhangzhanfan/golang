/*
	map遍历 删除  修改
*/

package main

import "fmt"

func main() {
	var countryCapitalMap map[string]string /*创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India"] = "新德里"

	//遍历
	for key, value := range countryCapitalMap {
		fmt.Println(key, "首都是", value)
	}
	fmt.Println("--------------------------")

	// 删除
	delete(countryCapitalMap, "India")

	// 修改
	countryCapitalMap["Japan"] = "xxx"

	for key, value := range countryCapitalMap {
		fmt.Println(key, "首都是", value)
	}

	// /*查看元素在集合中是否存在 */
	// capital, ok := countryCapitalMap["American"] /*如果确定是真实的,则存在,否则不存在 */
	// /*fmt.Println(capital) */
	// /*fmt.Println(ok) */
	// if ok {
	// 	fmt.Println("American 的首都是", capital)
	// } else {
	// 	fmt.Println("American 的首都不存在")
	// }
}
