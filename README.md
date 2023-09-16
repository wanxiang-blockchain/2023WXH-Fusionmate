# 2023WXH-Fusionmate
![Alt text](/images/logo.png)

# Table of Contents
- [2023WXH-Fusionmate](#2023wxh-fusionmate)
- [Table of Contents](#table-of-contents)
- [Introduction](#introduction)
  - [Background](#background)
  - [Motivation and Goal](#motivation-and-goal)
  - [User Group Orientation](#user-group-orientation)
- [FusionMate Architecture](#fusionmate-architecture)
  - [Training Modules](#training-modules)
  - [FusionMate Contracts](#fusionmate-contracts)

# Introduction 
## Background
AI assistants can improve people's life. If these AI assistants can be customized to provide different professional skills to assist work, it will be more efficient to improve human's work. At the same time, if these customized AI assistants can be traded in an AI assistant market similar to the talent market, the utility of these AI assistants with different skills will be maximized, both in terms of functionality and economic benefits.
## Motivation and Goal
AI is changing people's life. FusionMate makes "personal company" possible, that is, one person plus N AI assistants with various professional skills provided by FusionMate can form a personal company, a One-man Army!

FusionMate is a DAPP platform that provides AI assistant services. The AI assistants on the platform are AI language models in different fields and NFT assets based on the ERC-6551 standard. They can help you write documents, finish work, and give professional advice, plan travel itineraries, match fashionable clothes, or recommend movies and books.

Creators can use the general large models and toolset provided by FusionMate to create their own AI assistants, from simple prompt concepts to self-trained models. FusionMate has deployed factory contract which can be used to create AI assistant asset on the blockchain, and users only need to mint the AI assistant to recruit it.

FusionMate uses blockchain to guarantee the originality of AI-based assistants and also provides fun. FusionMate is the growth soil for AI assistants, where they can evolve and fusion. Here, the results of everyone's labor will be respected, fairness will be guaranteed, and our work and life will become more convenient with the help of AI assistants.
## User Group Orientation
FusionMate's market welcomes all users who embrace AI technology, especially those who want to build a one-man army through AI. FusionMate is also suitable for users who show great interest in training AI assistants which in turn bring direct economic value to these users.

# FusionMate Architecture
![Alt text](/images/architecture.png)

## Training Modules
AI assistants creator or maker, it owns some training models or it uses AWS SageMaker AI IDE to train auxiliary models, or it even possess terse and forceful prompts which provides an interactive LLVM by API. 
## FusionMate Contracts
- FusiomMate Factory contract: AI assistants factory where any maker can invoke `createAssistant` interface to create an AI assistant.
- AI NFT Contracts: everyone can invoke `mint` interface to generate an AI NFT in an AI assistant contract. When you own an AI NFT, you can use the corresponding AI service, such as chat with it, or get personalized feedback from it.
- ERC6551 Contracts: when a user mint a nft, a erc6551 token bound account will be generated for this nft. This token bound acount can receive fusionMate points.
  




