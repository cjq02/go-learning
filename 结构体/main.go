package main

import "fmt"

// ========== 基础结构体 ==========

// Animal 基础结构体（类似父类）
type Animal struct {
	Name string
	Age  int
}

// Animal 的方法
func (a *Animal) Eat() {
	fmt.Printf("%s 正在吃东西\n", a.Name)
}

func (a *Animal) Sleep() {
	fmt.Printf("%s 正在睡觉\n", a.Name)
}

func (a *Animal) GetInfo() {
	fmt.Printf("动物信息: 姓名=%s, 年龄=%d\n", a.Name, a.Age)
}

// ========== 通过嵌入实现"继承" ==========

// Dog 结构体嵌入 Animal（类似继承）
type Dog struct {
	Animal        // 嵌入 Animal，Dog 自动拥有 Animal 的所有字段和方法
	Breed  string // Dog 特有的字段
}

// Dog 特有的方法
func (d *Dog) Bark() {
	fmt.Printf("%s 正在汪汪叫\n", d.Name)
}

// Dog 可以重写（覆盖）父结构体的方法
func (d *Dog) GetInfo() {
	fmt.Printf("狗狗信息: 姓名=%s, 年龄=%d, 品种=%s\n", d.Name, d.Age, d.Breed)
}

// Cat 结构体嵌入 Animal
type Cat struct {
	Animal        // 嵌入 Animal
	Color  string // Cat 特有的字段
}

// Cat 特有的方法
func (c *Cat) Meow() {
	fmt.Printf("%s 正在喵喵叫\n", c.Name)
}

// ========== 多级嵌入 ==========

// Pet 结构体
type Pet struct {
	Owner string
}

func (p *Pet) GetOwner() {
	fmt.Printf("主人是: %s\n", p.Owner)
}

// GoldenRetriever 多重嵌入
type GoldenRetriever struct {
	Dog  // 嵌入 Dog（Dog 已经嵌入了 Animal）
	Pet  // 嵌入 Pet
	Size string
}

func (g *GoldenRetriever) GetInfo() {
	fmt.Printf("金毛信息: 姓名=%s, 年龄=%d, 品种=%s, 大小=%s, 主人=%s\n",
		g.Name, g.Age, g.Breed, g.Size, g.Owner)
}

func main() {
	fmt.Println("=== Go 语言结构体嵌入（类似继承）示例 ===")
	fmt.Println()

	// ========== 示例1: 基础嵌入 ==========
	fmt.Println("--- 示例1: Dog 嵌入 Animal ---")
	dog := Dog{
		Animal: Animal{
			Name: "旺财",
			Age:  3,
		},
		Breed: "哈士奇",
	}

	// Dog 可以直接使用 Animal 的字段和方法
	fmt.Println("直接访问嵌入的字段:", dog.Name)
	dog.Eat()     // 调用 Animal 的方法
	dog.Sleep()   // 调用 Animal 的方法
	dog.Bark()    // 调用 Dog 自己的方法
	dog.GetInfo() // 调用 Dog 重写的方法

	fmt.Println()

	// ========== 示例2: Cat 嵌入 Animal ==========
	fmt.Println("--- 示例2: Cat 嵌入 Animal ---")
	cat := Cat{
		Animal: Animal{
			Name: "小花",
			Age:  2,
		},
		Color: "橘色",
	}

	cat.Eat()     // 使用 Animal 的方法
	cat.Meow()    // 使用 Cat 自己的方法
	cat.GetInfo() // 如果没有重写，会使用 Animal 的 GetInfo

	fmt.Println()

	// ========== 示例3: 多重嵌入 ==========
	fmt.Println("--- 示例3: GoldenRetriever 多重嵌入 ---")
	golden := GoldenRetriever{
		Dog: Dog{
			Animal: Animal{
				Name: "金毛",
				Age:  5,
			},
			Breed: "金毛寻回犬",
		},
		Pet: Pet{
			Owner: "张三",
		},
		Size: "大型",
	}

	// 可以访问所有嵌入结构体的字段和方法
	golden.Eat()      // Animal 的方法
	golden.Bark()     // Dog 的方法
	golden.GetOwner() // Pet 的方法
	golden.GetInfo()  // GoldenRetriever 重写的方法

	fmt.Println()

	// ========== 示例4: 嵌入的匿名访问 ==========
	fmt.Println("--- 示例4: 嵌入的匿名访问和显式访问 ---")
	dog2 := Dog{
		Animal: Animal{Name: "小黑", Age: 4},
		Breed:  "拉布拉多",
	}

	// 匿名访问（推荐）
	fmt.Println("匿名访问 Name:", dog2.Name)
	dog2.Eat()

	// 显式访问（也可以这样写）
	fmt.Println("显式访问 Name:", dog2.Animal.Name)
	dog2.Animal.Eat()

	fmt.Println()

	// ========== 示例5: 方法集和方法提升 ==========
	fmt.Println("--- 示例5: 方法提升 ---")
	// 嵌入的结构体的方法会被"提升"到外层结构体
	// 可以直接在外层结构体上调用内层结构体的方法

	var animal Animal = Animal{Name: "通用动物", Age: 1}
	var dog3 Dog = Dog{
		Animal: Animal{Name: "测试狗", Age: 2},
		Breed:  "测试品种",
	}

	// 注意：虽然 Dog 嵌入了 Animal，但 Dog 和 Animal 是不同的类型
	// 不能直接将 Dog 赋值给 Animal（需要显式转换或使用接口）
	fmt.Println("Animal:", animal)
	fmt.Println("Dog:", dog3)
}
