erDiagram

authors {
	id id
	name name
}


book_authors {
	book_id book_id
	author_id author_id
}


book_categories {
	book_id book_id
	category_id category_id
}


books {
	id id
	title title
	published_date published_date
}


categories {
	id id
	category category
}

authors ||--|{ book_authors: ""
books ||--|{ book_categories: ""
books ||--|{ book_authors: ""
categories ||--|{ book_categories: ""
