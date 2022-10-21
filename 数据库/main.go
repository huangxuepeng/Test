package main

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Detail struct {
	ID           int    `gorm:"type:int;primaryKey"`
	City         string `gorm:"column:city;type:varchar(1000);not null;comment '城市'"`
	Compare      string `gorm:"column:compare;type:varchar(1000);comment '比较与演变'"`
	Genesis      string `gorm:"column:genesis;type:varchar(1000);comment '成因'"`
	Building     string `gorm:"column:building;type:varchar(1000);comment '建筑'"`
	Distribution string `gorm:"column:distribution;type:varchar(1000);comment '分布'"`
	Image        string `gorm:"column:image;type:varchar(1000);comment '图片'"`
	Build        string `gorm:"column:build;type:varchar(1000);comment '建造'"`
	Residence    string `gorm:"column:residence;type:varchar(1000);comment '民居'"`
	Type         string `gorm:"column:type;type:varchar(1000);comment '民居类型'"`
	Shape        string `gorm:"column:shape;type:varchar(1000);comment '形制'"`
	Owner        string `gorm:"column:owner;type:varchar(1000);comment '整理者'"`
}

var DB *gorm.DB

func init() {
	dsn := "root:12345678@tcp(119.45.221.151:3306)/city?charset=utf8mb4&parseTime=True&loc=Local"
	DB, _ = gorm.Open(mysql.Open(dsn))
	// DB.AutoMigrate(&Detail{})
}

func main() {
	aa := &Detail{
		City:         "澳门",
		Compare:      "碉楼式街屋由于功能特殊，多采用类似的前店后楼的形式。营业规模较大的当铺往往会将二、三层作为办公使用，而普通的当铺有时会沿袭传统街屋下铺上居的格局，作为私人的居所使用。",
		Genesis:      "碉楼式街屋是传统街屋的特殊形式，继承了传统防御性建筑的形式，来适应当铺对商业功能的独特要求。在当时平均建筑高度不超过十多米的城市中具有高度的视觉辨识性，也满足了当铺追求广告效果的要求。",
		Building:     "德生大按：位于澳门十月初五街65号，建于19纪。该建筑下层为当铺，当铺后有单独的高楼作为存贮典当物件的场所。存贮楼为五层，只在上部朝内院方向开设小窗，沿街侧面不开窗以确保安全。外墙用青砖砌筑，正面右侧上层采用了半圆券形窗，外面加有木白叶窗扇，是受到葡萄牙建筑形制影响的表现，反映出19世纪澳门中西式建筑的交融（图5）。",
		Distribution: "碉楼式街屋随着澳门19世纪末商业的繁盛而兴起，其主要分布在华人区商业活动最为繁荣的地段，如19世纪末的河边新街和20世纪的新马路，往往毗邻街角建造，避免窃贼利用相邻建筑攀入调楼，保证安全性。",
		Build:        "与一般的街屋相似，当铺以青砖墙体和木檩条承搭屋顶，建筑物屋顶为木结构，铺瓦，外墙上部有“某某大按”的灰塑字样。其特点在于碉楼的建造强调坚固与密封，因此多使用数米高的青石砌底，上层的青砖也比一般街屋厚数倍，如板樟堂街6号的高升大按。20世纪20年代建造的碉楼，则多使用最新的钢筋混凝土技术，上层墙面和底层柱墩用水刷石立面，如现存新马路的德成按当铺（图2、图3）。",
		Residence:    "碉楼式街屋是澳门中式街屋中最为特殊的一种，是专用于中国传统典当行业的建筑类型。其形式发展自传统街屋的“前铺后宅”形式，而把“后宅”改为“后楼”，以确保贵重物品的保存，在建造上注重防盗和私密性的功能需求。",
		Type:         "店铺式民居·碉楼式街屋",
		Shape:        "一间典型的前铺后（货）楼格局的当铺，分为前厅和主楼两个部分。入口有传统木趟门。其平面分为前后两厅，其间墙下为砌石，上为铁栏杆，顾客只能站在前厅与后厅之间议价。当楼高度一般三层，上层为办公室。后厅之后是天井，然后是碉楼式的货楼：货楼高度一般超过20m，平面呈正方形。底部为砌石结构，上为两层空心青砖墙。建筑物四面都设多个窄长的小窗，既防盗也利于通风。室内常常设置青砖墙分割空间来安放木货架（图4、图6）。",
		Owner:        "信息与控制工程学院2021级研究生席志国（学号：2108211581）",
	}
	DB.Create(&aa)
}
