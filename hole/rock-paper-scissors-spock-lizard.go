package hole

func rockPaperScissorsSpockLizard() []Scorecard {
	return outputTests(shuffle([]test{
		{"💎💎", "Tie"},
		{"💎📄", "📄 covers 💎"},
		{"💎✂", "💎 crushes ✂"},
		{"💎🖖", "🖖 vaporizes 💎"},
		{"💎🦎", "💎 crushes 🦎"},
		{"📄💎", "📄 covers 💎"},
		{"📄📄", "Tie"},
		{"📄✂", "✂ cuts 📄"},
		{"📄🖖", "📄 disproves 🖖"},
		{"📄🦎", "🦎 eats 📄"},
		{"✂💎", "💎 crushes ✂"},
		{"✂📄", "✂ cuts 📄"},
		{"✂✂", "Tie"},
		{"✂🖖", "🖖 smashes ✂"},
		{"✂🦎", "✂ decapitates 🦎"},
		{"🖖💎", "🖖 vaporizes 💎"},
		{"🖖📄", "📄 disproves 🖖"},
		{"🖖✂", "🖖 smashes ✂"},
		{"🖖🖖", "Tie"},
		{"🖖🦎", "🦎 poisons 🖖"},
		{"🦎💎", "💎 crushes 🦎"},
		{"🦎📄", "🦎 eats 📄"},
		{"🦎✂", "✂ decapitates 🦎"},
		{"🦎🖖", "🦎 poisons 🖖"},
		{"🦎🦎", "Tie"},
	}))
}
