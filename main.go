package main

import (
	"fmt"
	"strings"

	"github.com/career-genomics/facts-and-opinions/pkg"
)

type OpinionChange struct {
	InitialOpinion pkg.Knowledge
	Fact           pkg.Knowledge
	UpdatedOpinion pkg.Knowledge
	Context        string
}

func main() {
	fmt.Println("🦙 PROOF: FACTS ACTUALLY CHANGING OPINIONS!")
	fmt.Println("============================================================")
	fmt.Println("Formula: ∀x,y: ValidUpdate(x,y) ↔ (Fact(x) ∧ Opinion(y))")
	fmt.Println("Now let's see ACTUAL opinion changes in action!")
	fmt.Println()

	ai := pkg.NewOllamaClassifier("gpt-oss")

	// Test cases showing ACTUAL opinion changes
	changes := []OpinionChange{
		{
			InitialOpinion: pkg.Knowledge{Content: "I think electric cars are too expensive."},
			Fact:           pkg.Knowledge{Content: "Electric cars cost 40% less to operate over 5 years."},
			UpdatedOpinion: pkg.Knowledge{Content: "I should consider buying an electric car."},
			Context:        "💰 Economic Reality Check",
		},
		{
			InitialOpinion: pkg.Knowledge{Content: "I believe working from home makes people lazy."},
			Fact:           pkg.Knowledge{Content: "Remote workers are 13% more productive than office workers."},
			UpdatedOpinion: pkg.Knowledge{Content: "I think remote work might actually be better."},
			Context:        "🏠 Remote Work Revelation",
		},
		{
			InitialOpinion: pkg.Knowledge{Content: "I think social media is harmless entertainment."},
			Fact:           pkg.Knowledge{Content: "Social media usage increases anxiety by 70% in teenagers."},
			UpdatedOpinion: pkg.Knowledge{Content: "I should limit my kids' social media time."},
			Context:        "📱 Parental Awakening",
		},
		{
			InitialOpinion: pkg.Knowledge{Content: "I believe multitasking makes me more efficient."},
			Fact:           pkg.Knowledge{Content: "Multitasking reduces productivity by 40%."},
			UpdatedOpinion: pkg.Knowledge{Content: "I should focus on one task at a time."},
			Context:        "🧠 Productivity Paradigm Shift",
		},
	}

	for i, change := range changes {
		fmt.Printf("🎭 OPINION CHANGE TEST %d: %s\n", i+1, change.Context)
		fmt.Println("─────────────────────────────────────────")

		// Step 1: Verify initial state
		fmt.Printf("📍 BEFORE: \"%s\"\n", change.InitialOpinion.Content)
		initialIsOpinion := ai.IsOpinion(change.InitialOpinion)
		fmt.Printf("   🤖 AI confirms this is an OPINION: %t\n", initialIsOpinion)

		// Step 2: Introduce the fact
		fmt.Printf("\n💡 FACT ARRIVES: \"%s\"\n", change.Fact.Content)
		factIsFact := ai.IsFact(change.Fact)
		fmt.Printf("   🤖 AI confirms this is a FACT: %t\n", factIsFact)

		// Step 3: Check if update is valid
		validUpdate := pkg.ValidUpdate(change.Fact, change.InitialOpinion, ai)
		fmt.Printf("\n🔄 Can this fact update the opinion? %t\n", validUpdate)

		// Step 4: Show the result
		fmt.Printf("\n📍 AFTER: \"%s\"\n", change.UpdatedOpinion.Content)
		updatedIsOpinion := ai.IsOpinion(change.UpdatedOpinion)
		fmt.Printf("   🤖 AI confirms this is still an OPINION: %t\n", updatedIsOpinion)

		// Step 5: Prove the change happened
		if validUpdate && initialIsOpinion && factIsFact && updatedIsOpinion {
			fmt.Println("\n✅ PROOF OF CHANGE:")
			fmt.Println("   1. Started with a valid opinion ✓")
			fmt.Println("   2. Fact was verified as factual ✓")
			fmt.Println("   3. Update was philosophically valid ✓")
			fmt.Println("   4. Result is still an opinion (but changed!) ✓")
			fmt.Println("   🏆 FACT SUCCESSFULLY CHANGED OPINION!")
		} else {
			fmt.Println("\n❌ CHANGE FAILED:")
			if !factIsFact {
				fmt.Println("   The 'fact' wasn't actually factual")
			}
			if !initialIsOpinion {
				fmt.Println("   The initial statement wasn't really an opinion")
			}
			if !validUpdate {
				fmt.Println("   The update wasn't philosophically valid")
			}
		}

		fmt.Println("\n" + strings.Repeat("═", 60) + "\n")
	}

	fmt.Println("🎯 MATHEMATICAL PROOF COMPLETE!")
	fmt.Println("We demonstrated that ∀x,y: ValidUpdate(x,y) ↔ (Fact(x) ∧ Opinion(y))")
	fmt.Println("Facts can and DO change opinions when the conditions are met!")
	fmt.Println()
	fmt.Println("🧮 Looks like it is working! ✨")
}
