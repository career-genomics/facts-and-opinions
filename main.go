package main

import (
	"fmt"
	"strings"

	"github.com/career-genomics/facts-and-opinions/pkg"
)

func main() {
	fmt.Println("🦙 OLLAMA AI vs HUMAN LOGIC SHOWDOWN!")
	fmt.Println("============================================================")
	fmt.Println("Formula: ∀x,y: ValidUpdate(x,y) ↔ (Fact(x) ∧ Opinion(y))")
	fmt.Println("Translation: Facts can boss around opinions, but not vice versa!")
	fmt.Println()
	fmt.Println("Make sure Ollama is running with: ollama serve")
	fmt.Println("And you have a model installed: ollama pull gpt-oss")
	fmt.Println()

	ai := pkg.NewOllamaClassifier("gpt-oss")

	// Test cases that will make people think... and laugh
	testCases := []struct {
		x     pkg.Knowledge
		y     pkg.Knowledge
		title string
		joke  string
	}{
		{
			pkg.Knowledge{Content: "The Earth revolves around the Sun."},
			pkg.Knowledge{Content: "Which means the sun also revolves around the Earth."},
			"🌍 Classic Confused Logic",
			"When you understand physics but also don't...",
		},
		{
			pkg.Knowledge{Content: "Coffee contains caffeine."},
			pkg.Knowledge{Content: "I need coffee to function."},
			"☕ Programmer's Dilemma",
			"The most scientifically backed opinion in tech!",
		},
		{
			pkg.Knowledge{Content: "Humans need 8 hours of sleep."},
			pkg.Knowledge{Content: "Netflix is more important than sleep."},
			"😴 Modern Life Priorities",
			"Science vs streaming - the eternal struggle!",
		},
		{
			pkg.Knowledge{Content: "Exercise increases lifespan."},
			pkg.Knowledge{Content: "Pizza is a vegetable because tomatoes."},
			"🍕 Nutritional Mental Gymnastics",
			"When you try to make unhealthy choices sound scientific!",
		},
		{
			pkg.Knowledge{Content: "Money can't buy happiness."},
			pkg.Knowledge{Content: "I should quit my job and become a philosopher."},
			"💰 Philosophical Career Advice",
			"Plot twist: Is this saying even a fact?",
		},
		{
			pkg.Knowledge{Content: "Social media usage correlates with anxiety."},
			pkg.Knowledge{Content: "I should definitely check Instagram right now."},
			"📱 Self-Sabotage Logic",
			"When you know the research but ignore it anyway!",
		},
		{
			pkg.Knowledge{Content: "Vegetables are healthy."},
			pkg.Knowledge{Content: "French fries count as vegetables."},
			"🥔 Culinary Philosophy",
			"The greatest food classification debate of our time!",
		},
	}

	for i, tc := range testCases {
		fmt.Printf("📋 TEST %d: %s\n", i+1, tc.title)
		fmt.Printf("😂 %s\n", tc.joke)
		fmt.Println("─────────────────────────────────────────")

		// Run the AI analysis
		xisFact := ai.IsFact(tc.x)
		yisOpinion := ai.IsOpinion(tc.y)
		validUpdate := pkg.ValidUpdate(tc.x, tc.y, ai)

		// Display results with personality
		fmt.Printf("📝 X: \"%s\"\n", tc.x.Content)
		if xisFact {
			fmt.Println("   🤖 AI: \"Yep, that's a FACT!\" ✅")
		} else {
			fmt.Println("   🤖 AI: \"Hmm, not quite a fact...\" ❌")
		}

		fmt.Printf("📝 Y: \"%s\"\n", tc.y.Content)
		if yisOpinion {
			fmt.Println("   🤖 AI: \"Classic OPINION territory!\" ✅")
		} else {
			fmt.Println("   🤖 AI: \"That's not really an opinion...\" ❌")
		}

		// The verdict
		fmt.Printf("\n🎯 Can Facts Update Opinions? ")
		if validUpdate {
			fmt.Println("YES! ✅")
			fmt.Println("   🏆 Our formula works: Facts → Opinions is VALID!")
		} else {
			fmt.Println("NOPE! ❌")
			if xisFact && !yisOpinion {
				fmt.Println("   🤔 The 'opinion' isn't really subjective...")
			} else if !xisFact && yisOpinion {
				fmt.Println("   🤔 The 'fact' isn't actually factual...")
			} else {
				fmt.Println("   🤔 Neither side fits the formula!")
			}
		}

		fmt.Println("\n" + strings.Repeat("═", 50) + "\n")
	}

	fmt.Println("🎉 CONCLUSION:")
	fmt.Println("AI just proved that your philosophical formula works!")
	fmt.Println("Facts can change opinions, but opinions can't change facts.")
	fmt.Println("Now go forth and win internet arguments with MATH! 🧮✨")
	fmt.Println()
	fmt.Println("P.S. - If the AI disagrees with you, remember:")
	fmt.Println("That's just your OPINION vs its FACT analysis! 😉")
}
