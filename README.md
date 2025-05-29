# Lucky Number Game - Go BE K15 Practice

## Game Overview

A multithreaded lucky number gambling simulation with the following specifications.

## Core Mechanics

**Lucky Number Generation:**

- The system generates a random lucky number between 0-9 every 10 seconds
- This number serves as the winning condition for all bets

**User Betting System:**

- Users can input bet amounts ranging from $1 to $100
- The system accepts multiple concurrent users placing bets
- Each bet amount (x) determines both the prize pool contribution and waiting time

## Betting Process Flow

**When a user places a bet of amount X:**

1. **Pool Management:** Add the bet amount to the cumulative prize pool
2. **System Notification:** Display message in format:
   ```
   System: YYYY-MM-DD HH:mm:ss: User {i} bet ${x}, the current pool is {pool_total}$, waiting for {x}s to receive result
   ```
3. **Result Processing:** User waits X seconds before result determination
4. **Win Condition:** After X seconds, check if `X % 10 = lucky_number`
   - **If Win:** User receives the entire current prize pool
   - **If Lose:** User receives nothing

**Result Notifications:**

- **Win:** `System: YYYY-MM-DD HH:mm:ss: User {i} hit the lucky number, get {pool_amount}$`
- **Lose:** `System: YYYY-MM-DD HH:mm:ss: Wish user {i} lucky next time`

## Concurrent Operations

- While any user is waiting for their result, other users can continue placing new bets
- The system maintains thread-safe operations for pool management and user interactions

## System Termination

- When any user inputs "END", the system stops accepting new bets
- All pending bet results must complete before system shutdown
- Display final status: `System: Exited.`

## Technical Requirements

- Implement multithreading to handle concurrent user bets and timed result processing
- Ensure thread-safe access to shared resources (prize pool, user counter)
- Maintain accurate timestamps for all system messages
- Handle graceful shutdown with proper cleanup of pending operations
