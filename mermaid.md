erDiagram

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
	int author_id
	int book_id
}


books_categories {
	int category_id
	int book_id
}


categories {
	int id
	string category
}

authors ||--|{ books_authors: ""
books ||--|{ books_categories: ""
books ||--|{ books_authors: ""
categories ||--|{ books_categories: ""
