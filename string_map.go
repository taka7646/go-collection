package collection

import "sort"

// StringMap 文字列キーのMap
type StringMap struct {
	Items       map[string]interface{}
	valueSorter func(v1, v2 interface{}) bool
}

// NewStirngMap 新しいStringMapインスタンスを生成します
func NewStirngMap(valueSorter func(v1, v2 interface{}) bool) *StringMap {
	return &StringMap{
		Items:       make(map[string]interface{}),
		valueSorter: valueSorter,
	}
}

// isEmpty 空判定
func (m *StringMap) isEmpty() bool {
	if len(m.Items) == 0 {
		return true
	}
	return false
}

// Keys キー一覧を取得
func (m *StringMap) Keys() []string {
	keys := make([]string, 0, len(m.Items))
	for key := range m.Items {
		keys = append(keys, key)
	}
	return keys
}

// SortedKeys ソートずみキー一覧を取得
func (m *StringMap) SortedKeys() []string {
	keys := m.Keys()
	sort.Strings(keys)
	return keys
}

// Values 値の一覧を取得
func (m *StringMap) Values() []interface{} {
	values := make([]interface{}, 0, len(m.Items))
	for _, v := range m.Items {
		values = append(values, v)
	}
	if m.valueSorter != nil {
		sort.Slice(values, func(i, j int) bool {
			v1 := values[i]
			v2 := values[j]
			return m.valueSorter(v1, v2)
		})
	}
	return values
}

// HasKey キーの存在確認
func (m *StringMap) HasKey(key string) bool {
	_, ok := m.Items[key]
	return ok
}

// Set 値を設定
func (m *StringMap) Set(key string, value interface{}) {
	m.Items[key] = value
}

// Get デフォルト値つき取得
func (m *StringMap) Get(key string, defaultValue interface{}) interface{} {
	v, ok := m.Items[key]
	if ok {
		return v
	}
	return defaultValue
}

// MapM mutableなMapメソッド
// 全要素にmapFuncを適用して値を更新します
func (m *StringMap) MapM(mapFunc func(key string, value interface{}) interface{}) *StringMap {
	for k, v := range m.Items {
		m.Items[k] = mapFunc(k, v)
	}
	return m
}

// MapI immutableなMapメソッド
// 全要素にmapFuncを適用した新しいMapを返します
func (m *StringMap) MapI(mapFunc func(key string, value interface{}) interface{}) *StringMap {
	items := make(map[string]interface{})
	for k, v := range m.Items {
		items[k] = mapFunc(k, v)
	}
	return &StringMap{
		Items:       items,
		valueSorter: m.valueSorter,
	}
}

// FilterM mutableなFilterメソッド
// 全要素にmapFuncを適用して値を更新します
func (m *StringMap) FilterM(filterFunc func(key string, value interface{}) bool) *StringMap {
	for k, v := range m.Items {
		if !filterFunc(k, v) {
			delete(m.Items, k)
		}
	}
	return m
}

// FilterI immutableなFilterメソッド
// 全要素にfilterFuncを適用した新しいMapを返します
func (m *StringMap) FilterI(filterFunc func(key string, value interface{}) bool) *StringMap {
	items := make(map[string]interface{})
	for k, v := range m.Items {
		if filterFunc(k, v) {
			items[k] = v
		}
	}
	return &StringMap{
		Items:       items,
		valueSorter: m.valueSorter,
	}
}
