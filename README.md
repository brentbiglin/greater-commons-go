# Technology Learning Studio

* September
	* [September 1st](#2017-09-01) - Initial Post
	* [September 4th](#2017-09-04) - Description of Golang Motivations
	* [September 6th](#2017-09-06_01) - Topics, Goals, and Materials for Each of the 3 Projects
	* [September 6th](#2017-09-06_02) - Progress Report
	* [September 13th](#2017-09-13) - What the Func

## Learning Golang

### 2017-09-01

This readme acts as a blog for Professor Howison's Technology Learning Studio course at the University of Texas at Austin. I'll document my learning process here.

### 2017-09-04

My first learning project is to learn as much as I can of the Go programming language from the [greatercommons.com](greatercommons.com) learning site. The course is called "Learn How to Code: Google's Go (golang) Programming Language" and covers the basics of the language for true beginners of programming.

I've chosen this course because many resources and courses I've found for Go assume that one has no experience with the programming language in particular, but they tend to assume a background in programming or familiarity with another language like Java or C++. I have no background whatsoever in functional programming languages, so many of these would introduce the basics and then quickly lose me with exercises that would assume I knew how to solve a particular programming problem already and that the challenge would be to figure out how to implement it in Go.

My motivation for learning Go is to become familiar with programming languages in general so that I understand their potential and limitations. As a designer, I think this is very important. I also want to continually expand my skillset from a designer and researcher, to a front-end developer, to afull-stack designer and developer. Without a background in programming, I plan to pursue independent projects in programming as a hobby until I've (maybe) reached a point in the future when I could be paid to do it.

The reason I've chosen Go is because, from research I've done, it is considered a very straightforward language with a great emphasis on documentation and simple solutions. While there may be many ways to do something in another language, usually in Go there is only one or very few ways to accomplish something. These constraints, while diminishing freedom or creativity, lower the cognitive load for programmers and lowers the barrier for beginners. These observations are not my own, as I have only learned the basics of Go and have no programming background, but from research I've done online, people to whom I've reached out on Twitter in the Go community, and my software developer friend. Another reason I've decided to learn Go is because of the very engaged and welcoming Go community. The main forum is [https://forum.golangbridge.org/](https://forum.golangbridge.org/), and the main annual conference is [Gophercon](https://www.gophercon.com/). The two or three times that I've reached out to members of the Go community with questions about the language, they have always responded and reached out to their fellow community members to solicit their perspectives as well. In addition to those reasons, Go is very fast, supported by Google, and growing very quickly.

If for some reason I complete the course for Go, I will begin another course on the same Greater Commons site called, "Web Development with Google's Go Programming Language (golang)."

As I complete the course, my "Go Report Card" for the repository can be found [here](https://goreportcard.com/report/github.com/brentbiglin/greater-commons-go), which checks the format and "lint" of my exercise programs.

### 2017-09-06_01

#### Topic 1: Go

##### Goal

My simplest goal is to complete the "Learn How to Code: Google's Go (golang) Programming Language" course on the Greater Commons site. This course has 29 sections of about 5-7 lessons each. I believe it's a significant amount of work, but an achievable goal. If I complete this course, my stretch goal is to complete "Web Development with Google's Go Programming Language (golang)," which is also a Greater Commons course. It has 18 sections with the number of lessons ranging from 2 to 24. I am confident that these two courses provide more than enough work for one month.

##### Materials

All that is required is the Go programming language and the $GOPATH with a development environment. I'm using vim with a vim-go plugin, which allows me to format and run Go programs very easily. Each lesson is documented here in the repository.

#### Topic 2: Checkout System with Ponzu CMS and Server Framework

##### Goal

My goal is to create a checkout system or the various devices that are available to students from the IT Lab. Up until now, reservations, checkouts, and overdue devices have been recorded by hand. Presently, a workaround has been developed by the lab TAs to use Google Sheets and Google Forms. An optimized and standardized application developed for the lab's specific needs is required. The tool that I will use does not have a reservation/overdue notice feature, so my stretch goal will be developing and optimizing that..

##### Materials

I'll be using the [Ponzu CMS and Server Framework](https://docs.ponzu-cms.org/) to develop the application and requested features.

#### Topic 3: Drawing and Sketching

##### Goal

My goal is to make a noticeable improvement in my drawing and sketching abilities, particularly for sketching prototypes and user scenarios. I will sketch and draw a set group of figures/subjects and document my progress and improvement each day. My stretch goal is to become proficient in drawing human faces confidently.

##### Materials

I've purchased a set of [Uni Pin Drawing Pens](https://www.amazon.com/gp/product/B00F38T9C4/) with 6 assorted tip sizes and a [Strathmore Series 400](https://www.amazon.com/gp/product/B0027A39PY) sketch pad.

### 2017-09-06_02

#### Progress Update

I'm now finished with section 9 of 29 of the golang course. We've gone through the origins of Go and a bit about its creators; the fundaments of variables, values, and type; programming fundaments like bools, number types, string types, number systems like hexademical, binary, and decimal, constants, iota, and bit shifting; loops and conditionals; and arrays, slices, and maps. Some of the sections are exercise sections, in which I'm given a problem to solve or a program to write, and then I check my work. I had a bit of trouble with slices and arrays because of the syntax (too many brackets, curlies, and parentheses all together), but I plan on reviewing these sections. The next section, which deals primarily with structs, appears to build on these concepts so I think I'll continue to get more practice with them as I go along.

### 2017-09-13

#### What the Func

I've reached the part of my lesson where we're starting to deal with functions, their receivers, parameters, and returns. I've understood the basics pretty well, but I still feel like a lot of the concepts are uncharted territory. The lessons have definitely shifted gear into a bit more advanced concepts. One thing I had a lot of trouble with is the syntax for creating and referencing a slice.

```go
x := []int{1, 2, 3}
```
has the same effect as 

```go
x := make([]int, 3, 3)
x[0] = 1
x[1] = 2
x[2] = 3
```

It's just a few too many brackets, parentheses, and braces in different situations for me to get a handle on without a significant amount of practice. I think that this week I'll redo and review some of the exercises that are interspersed through the course so I don't move ahead before getting some of these fundamentals down, even though I doubt these things will become second-nature until/unless I start working on a project using Go in the future.

The Austin Go Meetup that I was planning on attending this month just released their topic, and it's "n00b Month." Very fortunate for me! They're going to go over the basics of the language, including the history, who invented it, and why. I'm hoping to meet other beginners in the language so that I can ask them how they're learning and what kind of resources they've found. I think that in general, it will be very beginner-friendly and many members will probably be expecting to field questions from people just starting out with the language.

I've also started to explore the official [Go Forum](https://forum.golangbridge.org/), which tends to have pretty advanced topics and very specific issues, and the [Gopher Academy Slack](https://blog.gopheracademy.com/gophers-slack-community/), which has a channel specifically for beginners.
