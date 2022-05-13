erDiagram

books_categories {
	int book_id
	int category_id
}


categories {
	int id
	string category
}


authors {
	int id
	string name
}


books {
	int id
	string title
	date published_date
}


books_authors {
	int book_id
	int author_id
}

authors ||--|{ books_authors: ""
books ||--|{ books_categories: ""
books ||--|{ books_authors: ""
categories ||--|{ books_categories: ""
