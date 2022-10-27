
package main


//expected data
type Post struct {
	Title, Description, Body string
	Tags                     []string
}

//what we want is  slicce of all posts in the folder
var posts []blogposts.Post


// make up a function to read in our data 
post = blogposts.NewPostsFromFS(someFS)