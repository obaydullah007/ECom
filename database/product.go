package database


var ProductList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}


func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Telescope",
		Description: "A Powerful Newtonian Telescope",
		Price:       199.99,
		ImgUrl:      "https://i.etsystatic.com/10189331/r/il/204377/3823588808/il_570xN.3823588808_cyr2.jpg",
	}
	prd2 := Product{
		ID:          2,
		Title:       "Bike",
		Description: "A Cafe Racer Bike",
		Price:       2999.99,
		ImgUrl:      "https://caferacerdreams.es/wp-content/uploads/2016/02/crd11-cafe-racer-crd-honda-cb750kz-portada.jpg",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Laptop",
		Description: "A High End Gaming Laptop",
		Price:       1499.99,
		ImgUrl:      "https://gagadget.com/media/cache/18/f9/18f934fbff1cdb6f3ff5949b1c3ab9c5.jpg",
	}
	prd4 := Product{
		ID:          4,
		Title:       "Smartphone",
		Description: "A Flagship Smartphone",
		Price:       999.99,
		ImgUrl:      "https://assets.gadgetandgear.com/upload/product/20240123_1705987892_224395.jpeg",
	}
	prd5 := Product{
		ID:          5,
		Title:       "Jeep",
		Description: "A Powerful Off-Road Jeep",
		Price:       49999.99,
		ImgUrl:      "https://5.imimg.com/data5/SELLER/Default/2020/9/PW/NC/HD/90810037/willy-open-jeep.jpg",
	}
	ProductList = append(ProductList, prd1, prd2, prd3, prd4, prd5)
}
