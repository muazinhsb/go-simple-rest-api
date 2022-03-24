package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

type Books []Book

func allBooks(w http.ResponseWriter, r *http.Request) {
	books := Books{
		Book{Title: "Harry Potter", Author: "J.K. Rowling", Desc: "tells the story of harry potter's life to become a great wizard who started from his miserable life"},
		Book{Title: "Percy Jackson and the Olympians", Author: "Rick Riordan", Desc: "the life of teenager Percy Jackson gets a lot more complicated when he learns he's the son of the Greek god Poseidon."},
		Book{Title: "The Lord of the Rings", Author: "J. R. R. Tolkien", Desc: "fate has placed it in the hands of a young Hobbit named Frodo Baggins, who inherits the Ring and steps into legend."},
		Book{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Desc: "Nick, a would-be writer, moves in next-door to millionaire Jay Gatsby and across the bay from his cousin Daisy and her philandering husband, Tom."},
		Book{Title: "Pride and Prejudice", Author: "Jane Austen", Desc: "Pride and Prejudice follows the turbulent relationship between Elizabeth Bennet, the daughter of a country gentleman, and Fitzwilliam Darcy, a rich aristocratic landowner."},
		Book{Title: "Death on the Nile", Author: "Agatha Christie", Desc: "Belgian sleuth Hercule Poirot's Egyptian vacation aboard a glamorous river steamer turns into a terrifying search for a murderer when a picture-perfect couple's idyllic honeymoon is tragically cut short."},
		Book{Title: "Goosebumps", Author: "R. L. Stine", Desc: "the stories follow child characters, who find themselves in scary situations, usually involving monsters and other supernatural elements."},
		Book{Title: "The Three Musketeers", Author: "Alexandre Dumas", Desc: "a historical romance, it relates the adventures of four fictional swashbuckling heroes who lived under the French kings Louis XIII and Louis XIV, who reigned during the 17th and early 18th centuries."},
		Book{Title: "Don Quixote", Author: "Miguel de Cervantes", Desc: "Don Quixote is a middle-aged gentleman from the region of La Mancha in central Spain. Obsessed with the chivalrous ideals touted in books he has read, he decides to take up his lance and sword to defend the helpless and destroy the wicked. "},
		Book{Title: "Da Vinci Code", Author: "Dan Brown", Desc: "The Da Vinci Code follows symbologist Robert Langdon and cryptologist Sophie Neveu after a murder in the Louvre Museum in Paris causes them to become involved in a battle between the Priory of Sion and Opus Dei over the possibility of Jesus Christ and Mary Magdalene having had a child together."},
	}

	fmt.Println("Endpoint Hit: All Books Endpoint")
	json.NewEncoder(w).Encode(books)
}

type Movie struct {
	Title    string `json:"title"`
	Director string `json:"director"`
	Desc     string `json:"desc"`
}

type Movies []Movie

func allMovies(w http.ResponseWriter, r *http.Request) {
	movies := Movies{
		Movie{Title: "Inception", Director: "Christopher Nolan", Desc: "Dom Cobb (Leonardo DiCaprio) is a thief with the rare ability to enter people's dreams and steal their secrets from their subconscious."},
		Movie{Title: "Shutter Island", Director: "Martin Scorsese", Desc: "the implausible escape of a brilliant murderess brings U.S. Marshal Teddy Daniels (Leonardo DiCaprio) and his new partner (Mark Ruffalo) to Ashecliffe Hospital, a fortress-like insane asylum located on a remote, windswept island."},
		Movie{Title: "The Shawshank Redemption", Director: "Frank Darabont", Desc: "Andy Dufresne (Tim Robbins) is sentenced to two consecutive life terms in prison for the murders of his wife and her lover and is sentenced to a tough prison. However, only Andy knows he didn't commit the crimes"},
		Movie{Title: "Forrest Gump", Director: "Robert Zemeckis", Desc: "slow-witted Forrest Gump (Tom Hanks) has never thought of himself as disadvantaged, and thanks to his supportive mother (Sally Field), he leads anything but a restricted life."},
		Movie{Title: "The Green Mile", Director: "Frank Darabont", Desc: "Paul Edgecomb (Tom Hanks) walked the mile with a variety of cons. He had never encountered someone like John Coffey (Michael Clarke Duncan), a massive black man convicted of brutally killing a pair of young sisters."},
		Movie{Title: "Saving Private Ryan", Director: "Steven Spielberg", Desc: "Captain John Miller (Tom Hanks) takes his men behind enemy lines to find Private James Ryan, whose three brothers have been killed in combat."},
		Movie{Title: "Scarface", Director: "Brian De Palma", Desc: "after getting a green card in exchange for assassinating a Cuban government official, Tony Montana (Al Pacino) stakes a claim on the drug trade in Miami."},
		Movie{Title: "Goodfellas", Director: "Martin Scorsese", Desc: "a young man grows up in the mob and works very hard to advance himself through the ranks. He enjoys his life of money and luxury, but is oblivious to the horror that he causes."},
		Movie{Title: "Taxi Driver", Director: "Martin Scorsese", Desc: "suffering from insomnia, disturbed loner Travis Bickle (Robert De Niro) takes a job as a New York City cabbie, haunting the streets nightly, growing increasingly detached from reality as he dreams of cleaning up the filthy city."},
		Movie{Title: "Interstellar", Director: "Christopher Nolan", Desc: "in Earth's future, a global crop blight and second Dust Bowl are slowly rendering the planet uninhabitable. Professor Brand (Michael Caine), a brilliant NASA physicist, is working on plans to save mankind by transporting Earth's population to a new home via a wormhole."},
	}

	fmt.Println("Endpoint Hit: All Movies Endpoint")
	json.NewEncoder(w).Encode(movies)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/books", allBooks)
	http.HandleFunc("/movies", allMovies)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequest()
}
