package collection

import "sort"

type Array struct {
	Items []interface{}
}

// Append 値の追加
func (m *Array) Append(value interface{}) {
	m.Items = append(m.Items, value)
}

// Sort ソート
func (m *Array) Sort(sortFunc func(v0, v1 interface{}) bool) {
	sort.Slice(m.Items, func(i, j int) bool {
		l := m.Items[i]
		r := m.Items[j]
		return sortFunc(l, r)
	})
}

// MapM mutableなMapメソッド
// 全要素にmapFuncを適用して値を更新します
func (m *Array) MapM(mapFunc func(i int, value interface{}) interface{}) *Array {
	for i, v := range m.Items {
		m.Items[i] = mapFunc(i, v)
	}
	return m
}

// MapI immutableなMapメソッド
// 全要素にmapFuncを適用した新しいMapを返します
func (m *Array) MapI(mapFunc func(i int, value interface{}) interface{}) *Array {
	items := make([]interface{}, 0, len(m.Items))
	for i, v := range m.Items {
		items[i] = mapFunc(i, v)
	}
	return &Array{
		Items: items,
	}
}

// FilterM mutableなFilterメソッド
// 全要素にmapFuncを適用して値を更新します
func (m *Array) FilterM(filterFunc func(i int, value interface{}) bool) *Array {
	items := make([]interface{}, 0, len(m.Items))
	for i, v := range m.Items {
		if filterFunc(i, v) {
			items = append(items, v)
		}
	}
	m.Items = items
	return m
}

// FilterI immutableなFilterメソッド
// 全要素にfilterFuncを適用した新しいMapを返します
func (m *Array) FilterI(filterFunc func(i int, value interface{}) bool) *Array {
	items := make([]interface{}, 0, len(m.Items))
	for i, v := range m.Items {
		if filterFunc(i, v) {
			items = append(items, v)
		}
	}
	return &Array{
		Items: items,
	}
}
