package main

import (
	"html/template"
	"net/http"
	"strconv"
)

// Structure produit
type Product struct {
	ID          int
	Name        string
	Price       int
	OldPrice    int
	Description string
	Image       string
	Sizes       string
	Discount    int
	Stock       int
}

var products = []Product{
	{ID: 1, Name: "PALACE PULL À CAPUCHE UNISEXE CHASSEUR", Price: 148, OldPrice: 0, Description: "Un hoodie polyvalent conçu pour affronter le froid avec style. Sa coupe unisexe et sa capuche ajustable en font une pièce confortable et pratique. Fabriqué dans un coton épais et doux, il offre chaleur et durabilité. Parfait pour un look urbain affirmé.", Image: "/products/hoodie1.webp", Sizes: "S | M | L", Discount: 0, Stock: 12},
	{ID: 2, Name: "PALACE PULL À CAPUCHON MARINE", Price: 138, OldPrice: 0, Description: "Confectionné dans un tissu doux et confortable, ce hoodie marine est un incontournable de la saison. Sa couleur sobre et élégante s’associe facilement à toutes les tenues. Le logo Palace brodé sur la poitrine apporte une touche streetwear authentique.", Image: "/products/hoodie2.webp", Sizes: "S | XS | M", Discount: 0, Stock: 5},
	{ID: 3, Name: "PALACE PULL CREW PASSEPOIL NOIR", Price: 115, OldPrice: 145, Description: "Un sweat noir classique revisité avec des détails passepoilés discrets. Sa coupe droite et son tissu molletonné en font un vêtement agréable à porter au quotidien. Idéal pour un style minimaliste mais affirmé, il s’adapte aussi bien à un jean qu’à un pantalon cargo.", Image: "/products/sweetshirt.webp", Sizes: "M | L", Discount: 20, Stock: 0},
	{ID: 4, Name: "PALACE WASHED TERRY 1/4 PLACKET HOOD MOJITO", Price: 168, OldPrice: 0, Description: "Un hoodie original en coton éponge lavé, dans une teinte mojito rafraîchissante. Sa patte de boutonnage 1/4 ajoute une touche vintage et sportive. Léger mais résistant, il est parfait pour les journées plus douces ou pour superposer avec une veste.", Image: "/products/hoodie3.webp", Sizes: "S | M | L", Discount: 0, Stock: 8},
	{ID: 5, Name: "PALACE PANTALON BOSSY JEAN STONE", Price: 125, OldPrice: 0, Description: "Jean stone coupe droite, pensé pour un confort optimal et une allure décontractée. Sa toile robuste et ses finitions soignées garantissent une grande durabilité. Un basique intemporel qui s’associe facilement avec n’importe quel haut.", Image: "/products/jeans1.webp", Sizes: "32 | 34 | 36", Discount: 0, Stock: 15},
	{ID: 6, Name: "PALACE PANTALON CARGO GORE-TEX NOIR", Price: 110, OldPrice: 0, Description: "Un pantalon cargo technique en Gore-Tex, conçu pour résister aux intempéries. Ses poches multiples et sa coupe fonctionnelle en font un allié pratique pour la vie urbaine comme pour les sorties en extérieur. Allie performance et style streetwear.", Image: "/products/cargo.webp", Sizes: "M | L", Discount: 0, Stock: 3},
	{ID: 7, Name: "PALACE HOODIE CLASSIQUE", Price: 135, OldPrice: 0, Description: "Un hoodie simple et efficace, pensé pour un usage quotidien. Sa matière douce et respirante assure un confort optimal, tandis que sa coupe moderne s’adapte à toutes les morphologies. Un essentiel du vestiaire streetwear.", Image: "/products/hoodie4.webp", Sizes: "S | M | L", Discount: 0, Stock: 10},
}

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/product", productHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/add/submit", addSubmitHandler)

	http.Handle("/Css/", http.StripPrefix("/Css/", http.FileServer(http.Dir("Css"))))
	http.Handle("/products/", http.StripPrefix("/products/", http.FileServer(http.Dir("products"))))

	println("Serveur lancé sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", products)
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	for _, p := range products {
		if p.ID == id {
			tmpl.ExecuteTemplate(w, "product.html", p)
			return
		}
	}
	http.NotFound(w, r)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "add.html", nil)
}

func addSubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/add", http.StatusSeeOther)
		return
	}

	name := r.FormValue("name")
	priceStr := r.FormValue("price")
	desc := r.FormValue("description")
	sizes := r.FormValue("sizes")
	image := r.FormValue("image")
	stockStr := r.FormValue("stock")

	if name == "" || priceStr == "" || desc == "" || image == "" || stockStr == "" {
		http.Redirect(w, r, "/add", http.StatusSeeOther)
		return
	}

	price, _ := strconv.Atoi(priceStr)
	stock, _ := strconv.Atoi(stockStr)

	newProduct := Product{
		ID:          len(products) + 1,
		Name:        name,
		Price:       price,
		OldPrice:    0,
		Description: desc,
		Image:       "/products/" + image,
		Sizes:       sizes,
		Discount:    0,
		Stock:       stock,
	}

	products = append(products, newProduct)

	http.Redirect(w, r, "/product?id="+strconv.Itoa(newProduct.ID), http.StatusSeeOther)
}
