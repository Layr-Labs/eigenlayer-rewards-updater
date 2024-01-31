import { newMockEvent } from "matchstick-as"
import { ethereum, Address, BigInt } from "@graphprotocol/graph-ts"
import {
  Deposit,
  StrategyAddedToDepositWhitelist,
  StrategyRemovedFromDepositWhitelist,
  StrategyWhitelisterChanged
} from "../generated/StrategyManager/StrategyManager"

export function createDepositEvent(
  staker: Address,
  token: Address,
  strategy: Address,
  shares: BigInt
): Deposit {
  let depositEvent = changetype<Deposit>(newMockEvent())

  depositEvent.parameters = new Array()

  depositEvent.parameters.push(
    new ethereum.EventParam("staker", ethereum.Value.fromAddress(staker))
  )
  depositEvent.parameters.push(
    new ethereum.EventParam("token", ethereum.Value.fromAddress(token))
  )
  depositEvent.parameters.push(
    new ethereum.EventParam("strategy", ethereum.Value.fromAddress(strategy))
  )
  depositEvent.parameters.push(
    new ethereum.EventParam("shares", ethereum.Value.fromUnsignedBigInt(shares))
  )

  return depositEvent
}

export function createStrategyAddedToDepositWhitelistEvent(
  strategy: Address
): StrategyAddedToDepositWhitelist {
  let strategyAddedToDepositWhitelistEvent =
    changetype<StrategyAddedToDepositWhitelist>(newMockEvent())

  strategyAddedToDepositWhitelistEvent.parameters = new Array()

  strategyAddedToDepositWhitelistEvent.parameters.push(
    new ethereum.EventParam("strategy", ethereum.Value.fromAddress(strategy))
  )

  return strategyAddedToDepositWhitelistEvent
}

export function createStrategyRemovedFromDepositWhitelistEvent(
  strategy: Address
): StrategyRemovedFromDepositWhitelist {
  let strategyRemovedFromDepositWhitelistEvent =
    changetype<StrategyRemovedFromDepositWhitelist>(newMockEvent())

  strategyRemovedFromDepositWhitelistEvent.parameters = new Array()

  strategyRemovedFromDepositWhitelistEvent.parameters.push(
    new ethereum.EventParam("strategy", ethereum.Value.fromAddress(strategy))
  )

  return strategyRemovedFromDepositWhitelistEvent
}

export function createStrategyWhitelisterChangedEvent(
  previousAddress: Address,
  newAddress: Address
): StrategyWhitelisterChanged {
  let strategyWhitelisterChangedEvent = changetype<StrategyWhitelisterChanged>(
    newMockEvent()
  )

  strategyWhitelisterChangedEvent.parameters = new Array()

  strategyWhitelisterChangedEvent.parameters.push(
    new ethereum.EventParam(
      "previousAddress",
      ethereum.Value.fromAddress(previousAddress)
    )
  )
  strategyWhitelisterChangedEvent.parameters.push(
    new ethereum.EventParam(
      "newAddress",
      ethereum.Value.fromAddress(newAddress)
    )
  )

  return strategyWhitelisterChangedEvent
}
