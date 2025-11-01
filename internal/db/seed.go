package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"

	"github.com/Damisicode/social/internal/store"
)

var usernames = []string{
	"silverhawk22", "luna_coder", "techwave91", "nova_strider", "urbanleaf", "pixelpilot", "dataweaver", "nightbyte", "sonicflare", "mellowstone", "orbitdreamer", "cypheron", "codebender", "hazylogic", "frostengine", "cosmictrail", "midnightstack", "bytehunter", "ocean_rift", "echoatlas", "terraflux", "neonvoyage", "solarveil", "quantumleaf", "juno_trace", "skyflicker", "driftlogic", "vaporroot", "glitchmind", "zenorbit", "dusty_pixel", "brightstorm", "emberpath", "hexrider", "moonhopper", "nodechaser", "codeatlas", "deepnova", "silentloop", "rustveil", "microflux", "lightcascade", "stormweave", "datadream", "bitwanderer", "echoverse", "flowcraft", "binaryowl", "zencode", "synthtrail",
}

var titles = []string{
	"building scalable web apps", "intro to golang basics", "mastering clean architecture", "getting started with iot", "understanding cloud deployment", "optimizing api performance", "designing responsive interfaces", "secure authentication methods", "real time data streaming", "automating ci cd pipelines", "building smart waste systems", "introduction to nextjs", "using prisma with nestjs", "implementing jwt authentication", "creating reusable react components", "debugging race conditions go", "designing efficient microservices", "deploying apps on aws", "learning tailwind css", "modern software engineering practices",
}

var contents = []string{
	"how to build efficient backend systems",
	"getting started with embedded device programming",
	"understanding api design and versioning",
	"best practices for scalable web development",
	"introduction to internet of things concepts",
	"deploying full stack apps on aws",
	"using prisma orm with nestjs backend",
	"building real time dashboards with react",
	"how to secure restful api endpoints",
	"optimizing frontend performance with caching",
	"creating responsive ui with tailwind css",
	"understanding mqtt protocol for iot",
	"debugging and testing golang applications",
	"implementing role based access control",
	"automating deployment using github actions",
	"building communication between microservices",
	"how to use redis for caching",
	"managing authentication with magic links",
	"setting up private lorawan gateway",
	"integrating paystack payments in your app",
}

var tags = []string{
	"technology", "golang", "nestjs", "nextjs", "react", "iot", "cloud", "aws", "python", "javascript",
	"typescript", "frontend", "backend", "devops", "api", "microservices", "database", "prisma", "tailwind", "flask",
	"cicd", "github", "monitoring", "automation", "firmware", "hardware", "engineering", "innovation", "electronics", "wireless",
	"lorawan", "mqtt", "sensors", "datascience", "ai", "machinelearning", "security", "authentication", "networking", "deployment",
	"webdev", "uiux", "open_source", "productivity", "databases", "testing", "containers", "scalability", "architecture", "opensource",
}

var comments = []string{
	"great post, really informative!",
	"this explanation made things so clear.",
	"thanks for sharing such useful insights.",
	"i learned something new today.",
	"awesome breakdown of a complex topic.",
	"can you share the source code example?",
	"this helped me fix a similar issue.",
	"excellent write-up, keep it up!",
	"your posts are always so helpful.",
	"i love how you simplified this concept.",
	"very practical and easy to follow.",
	"looking forward to more content like this.",
	"this is exactly what i was searching for.",
	"the step-by-step guide was spot on.",
	"please make a follow-up article soon.",
	"clear, concise, and well explained.",
	"you saved me a lot of debugging time.",
	"thanks for sharing your experience!",
	"this topic deserves more attention.",
	"great work, i’m implementing this now.",
}

func Seed(store store.Storage, db *sql.DB) {
	ctx := context.Background()

	users := generateUsers(100)
	tx, _ := db.BeginTx(ctx, nil)

	for _, user := range users {
		if err := store.Users.Create(ctx, tx, user); err != nil {
			_ = tx.Rollback()
			log.Println("Error creating user:", err)
			return
		}
	}

	tx.Commit()
	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post:", err)
			return
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating comment:", err)
			return
		}
	}

	log.Println("Seeding complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			Role: store.Role{
				Name: "user",
			},
		}

		if err := users[i].Password.Set("123123"); err != nil {
			log.Println("Error setting password for user", users[i].Username)
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)
	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)
	for i := 0; i < num; i++ {
		cms[i] = &store.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  users[rand.Intn(len(users))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}

	return cms
}
