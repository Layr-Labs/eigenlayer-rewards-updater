import {
  Deposit as DepositEvent,
  StrategyAddedToDepositWhitelist as StrategyAddedToDepositWhitelistEvent,
  StrategyRemovedFromDepositWhitelist as StrategyRemovedFromDepositWhitelistEvent,
  StrategyWhitelisterChanged as StrategyWhitelisterChangedEvent
} from "../generated/StrategyManager/StrategyManager"
import {
  Deposit,
  StrategyAddedToDepositWhitelist,
  StrategyRemovedFromDepositWhitelist,
  StrategyWhitelisterChanged
} from "../generated/schema"

export function handleDeposit(event: DepositEvent): void {
  let entity = new Deposit(
    event.transaction.hash.concatI32(event.logIndex.toI32())
  )
  entity.staker = event.params.staker
  entity.token = event.params.token
  entity.strategy = event.params.strategy
  entity.shares = event.params.shares

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleStrategyAddedToDepositWhitelist(
  event: StrategyAddedToDepositWhitelistEvent
): void {
  let entity = new StrategyAddedToDepositWhitelist(
    event.transaction.hash.concatI32(event.logIndex.toI32())
  )
  entity.strategy = event.params.strategy

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleStrategyRemovedFromDepositWhitelist(
  event: StrategyRemovedFromDepositWhitelistEvent
): void {
  let entity = new StrategyRemovedFromDepositWhitelist(
    event.transaction.hash.concatI32(event.logIndex.toI32())
  )
  entity.strategy = event.params.strategy

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleStrategyWhitelisterChanged(
  event: StrategyWhitelisterChangedEvent
): void {
  let entity = new StrategyWhitelisterChanged(
    event.transaction.hash.concatI32(event.logIndex.toI32())
  )
  entity.previousAddress = event.params.previousAddress
  entity.newAddress = event.params.newAddress

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}
