package model

import (
	"fmt"
	"ginshop.com/utils"
	"gorm.io/gorm"
	"sort"
)

type Category struct {
	Model
	Name       string         `form:"name" binding:"required"  json:"name" gorm:"varchar(255);not null;default:'';comment:'分类名称'"  comment:"分类名称"`                    // 分类名称
	Pid        int            `form:"pid" json:"pid" gorm:"size:11;not null;default:0;comment:'分类节点'"`                                                                // 分类节点：0根节点
	Icon       string         `form:"icon" json:"icon" gorm:"varchar(255);not null;default:'';comment:'分类图标'" edit:"0"`                                               // 分类图标
	State      *int           `form:"state" json:"state" gorm:"size:1;not null;default:0;comment:'未启用,已启用'"  comment:"分类状态" text:"未启用,已启用" range:"0,1" types:"radio"` // 分类状态：0未启用，1已启用
	Sort       int64          `form:"sort" json:"sort" gorm:"size:11;not null;default:0;comment:'分类排序'" comment:"排序"`                                                 // 分类排序
	Tag        string         `form:"tag" json:"tag" gorm:"varchar(255);not null;default:0;comment:'分类标签'" edit:"0"`                                                  // 分类标签
	Children   *CategoryTrees `json:"children" gorm:"-"`
	InnerOrder int            `gorm:"inner_order" json:"-"`
	Level      int            `gorm:"level" json:"-"`
}

// CategoryTrees 二叉树列表
type CategoryTrees []*Category

// ToTree 转换为树形结构
func (Category) ToTree(data CategoryTrees, language string) CategoryTrees {
	// 定义 HashMap 的变量，并初始化
	TreeData := make(map[int]*Category)
	// 先重组数据：以数据的ID作为外层的key编号，以便下面进行子树的数据组合
	for _, item := range data {
		TreeData[item.Id] = item
	}

	// 定义 RoleTrees 结构体
	var TreeDataList CategoryTrees
	// 开始生成树形
	for _, item := range TreeData {
		// 如果没有根节点树，则为根节点
		if item.Pid == 0 {
			// 追加到 TreeDataList 结构体中
			item.Name = utils.Languagebycode(language, item.Name)
			TreeDataList = append(TreeDataList, item)
			// 跳过该次循环
			continue
		}
		item.Name = utils.Languagebycode(language, item.Name)
		// 通过 上面的 TreeData HashMap的组合，进行判断是否存在根节点
		// 如果存在根节点，则对应该节点进行处理
		if pItem, ok := TreeData[item.Pid]; ok {
			fmt.Println(pItem.Id)
			// 判断当次循环是否存在子节点，如果没有则作为子节点进行组合
			if pItem.Children == nil {
				// 写入子节点
				children := CategoryTrees{item}
				// 插入到 当次结构体的子节点字段中，以指针的方式
				pItem.Children = &children
				pItem.Name = utils.Languagebycode(language, pItem.Name)
				// 跳过当前循环
				continue
			}

			// 以指针地址的形式进行追加到结构体中
			*pItem.Children = append(*pItem.Children, item)
		}

	}
	//sort.Slice(TreeDataList, func(i, j int) bool {
	//	return TreeDataList[i].Id < TreeDataList[j].Id
	//})
	//SortCategoriesByID(TreeDataList)
	// 2. 递归处理子节点
	sort.Slice(TreeDataList, func(i, j int) bool {
		return TreeDataList[i].Id < TreeDataList[j].Id
	})
	return TreeDataList
}

func SortCategoriesByID(categories CategoryTrees) {
	if categories != nil {
		return
	}

	// 1. 对当前层级进行排序
	sort.Slice(categories, func(i, j int) bool {
		return categories[i].Id < categories[j].Id
	})

	// 2. 递归处理子节点
	for _, cat := range categories {
		if *cat.Children != nil {
			SortCategoriesByID(*cat.Children)
		}
	}
}

// GetCategoryIds 获取指定分类下的所有子分类编号
func (Category) GetCategoryIds(DB *gorm.DB, cateId int) (pids map[int]int) {
	if cateId == 0 {
		return
	}
	var Result []*Category
	DB.Select([]string{"id", "pid"}).Find(&Result)
	index := 0
	pids = make(map[int]int)
	// 递归遍历指定分类下的全部子分类编号
	var inPids func(Result []*Category, cateId int)
	inPids = func(Result []*Category, cateId int) {
		if Result == nil {
			return
		}
		for _, item := range Result {
			if item.Pid == cateId {
				pids[index] = item.Id
				index++
				inPids(Result, item.Id)
			}
		}
	}
	// 初始化
	inPids(Result, cateId)
	return
}
