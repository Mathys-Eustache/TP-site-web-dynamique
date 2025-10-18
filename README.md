StreetShop – TP Rendre un site web dynamique
🎯 Objectif du projet
Ce projet a été réalisé dans le cadre du TP « Rendre un site web dynamique » en Go (Golang).
L’objectif est de mettre en pratique la gestion des routes, des templates et des formulaires pour construire un mini site e‑commerce dynamique.

🚀 Fonctionnalités implémentées
Challenge 01 – Affichage de la liste des articles
- Page d’accueil / avec :
- Un en‑tête (logo + menu de navigation).
- Une liste de produits affichés dynamiquement depuis une variable globale.
- Pour chaque produit : image, nom, prix, réduction optionnelle.
- Un bouton Consulter qui redirige vers la page détail.
Challenge 02 – Page produit
- Route /product?id=... qui affiche les détails du produit sélectionné :
- Nom, description, image.
- Prix avec réduction si applicable.
- Tailles disponibles.
- Stock (affiché en vert si disponible, en rouge si rupture).
- Bouton Ajouter au panier (simulation).
Challenge 03 – Ajouter un produit
- Route /add qui affiche un formulaire permettant d’ajouter un produit :
- Nom, prix, description, tailles.
- Sélecteur d’image (parmi les fichiers disponibles).
- Champ stock disponible.
- Route /add/submit qui traite le formulaire :
- Vérifie les champs.
- Ajoute le produit à la liste globale.
- Redirige vers la page détail du nouveau produit.

📂 Organisation du projet
- main.go → logique serveur (routes, handlers, données).
- templates/ → fichiers HTML (index.html, product.html, add.html).
- Css/styles.css → mise en forme des pages.
- products/ → images des produits.
- logo/ → logo du site.

- Lancer le serveur :
go run main.go
- Ouvrir le site dans le navigateur :
👉 http://localhost:8080

✅ Points respectés
- Utilisation de routes distinctes (/, /product, /add, /add/submit).
- Utilisation de templates dynamiques avec boucles et conditions.
- Gestion des formulaires et ajout dynamique de produits.
- Mise en forme avec CSS pour un rendu proche des maquettes fournies.
- Ajout d’une fonctionnalité supplémentaire : gestion du stock.

👤 Auteur
Projet réalisé par Mathys – TP Golang Web.
