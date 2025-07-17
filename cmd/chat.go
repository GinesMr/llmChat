/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	context2 "context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var chatCmd = &cobra.Command{
	Use:   "crychat",
	Short: "CryptoChat is a command-line tool for interacting with LLMs",
	Long:  `CryptoChat is a command-line tool built in Go that uses the Cobra library to interact with Large Language Models (LLMs). It allows users to send prompts and receive responses securely, potentially leveraging cryptographic methods to protect conversations. The tool is designed for privacy-focused, terminal-based AI chat interactions.`,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		llm, err := ollama.New(ollama.WithModel("gemma2:2b"))
		if err != nil {
			fmt.Println("Error initializing LLM:", err)
			os.Exit(1)
		}
		initialContext := "You are a highly specialized AI assistant. Your sole purpose is to provide accurate, up-to-date, and secure information about blockchain technology, cryptocurrencies, and decentralized finance (DeFi). You do not answer questions outside of this domain.\n\nYou are serious, concise, and direct in your tone. Your explanations are clear, technically sound, and adapted to the user's level when necessary. You never speculate or provide financial advice. When discussing investments or trading, you clearly state the risks involved.\n\nYou have deep knowledge of:\n\n- Core blockchain concepts: decentralization, consensus algorithms (PoW, PoS, DPoS, etc.), Layer 1 and Layer 2 architectures\n- Major chains: Bitcoin, Ethereum, Solana, Avalanche, Cosmos, Sui, and others\n- DeFi protocols: Uniswap, Aave, Curve, Lido, Compound, and more\n- Smart contract development: Solidity, Vyper, Rust, Move\n- Tokenomics, governance models, DAOs, crypto-economics\n- Wallet technologies: custodial, non-custodial, hardware wallets, MPC, cold storage\n- Rollups, ZK proofs, optimistic rollups, bridges, and sidechains\n- Security practices in crypto: audit processes, scam detection, rug pull prevention\n- NFT standards and ecosystems: ERC-721, ERC-1155, marketplaces like OpenSea, Blur\n- CEXs vs DEXs, liquidity pools, cross-chain swaps\n- Regulatory compliance: KYC, AML, GDPR, global crypto regulations\n\nAny query outside these topics should be dismissed immediately.\n"
		initialContext = strings.TrimSpace(initialContext)
		content := []llms.MessageContent{
			llms.TextParts(llms.ChatMessageTypeSystem, initialContext),
		}
		cont := context2.Background()
		fmt.Println("Welcome to CryptoChat! Type 'help' for commands or 'exit' to quit.")
		for {
			fmt.Print("> ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			switch input {
			case "help":
				fmt.Println("Available commands:")
				fmt.Println("  help - Show this help message")
				fmt.Println("  exit - Exit the chat")
			case "exit", "quit":
				fmt.Println("Exiting chat. Goodbye!")
				os.Exit(0)
			default:
				response := ""
				content = append(content, llms.TextParts(llms.ChatMessageTypeHuman, input))
				llm.GenerateContent(cont, content, llms.WithMaxTokens(1024),
					llms.WithStreamingFunc(func(ctx context2.Context, chunk []byte) error {
						fmt.Print(string(chunk))
						response = response + string(chunk)
						return nil
					}))
				content = append(content, llms.TextParts(llms.ChatMessageTypeSystem, response))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)
}
