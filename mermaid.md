erDiagram

authors {
	int id
	varchar(45) name
}


book_authors {
	int book_id
	int author_id
}


book_categories {
	int book_id
	int category_id
}


books {
	int id
	varchar(60) title
	date published_date
}


categories {
	int id
	varchar(60) category
}

authors ||--|{ book_authors: ""
books ||--|{ book_categories: ""
books ||--|{ book_authors: ""
categories ||--|{ book_categories: ""
