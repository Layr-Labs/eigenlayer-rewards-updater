import { newMockEvent } from "matchstick-as"
import { ethereum, BigInt, Address, Bytes } from "@graphprotocol/graph-ts"
import {
  ActivationDelaySet,
  ClaimerSet,
  GlobalCommissionBipsSet,
  PaymentClaimed,
  PaymentUpdaterSet,
  RootSubmitted
} from "../generated/ClaimingManager/ClaimingManager"

export function createActivationDelaySetEvent(
  oldActivationDelay: BigInt,
  newActivationDelay: BigInt
): ActivationDelaySet {
  let activationDelaySetEvent = changetype<ActivationDelaySet>(newMockEvent())

  activationDelaySetEvent.parameters = new Array()

  activationDelaySetEvent.parameters.push(
    new ethereum.EventParam(
      "oldActivationDelay",
      ethereum.Value.fromUnsignedBigInt(oldActivationDelay)
    )
  )
  activationDelaySetEvent.parameters.push(
    new ethereum.EventParam(
      "newActivationDelay",
      ethereum.Value.fromUnsignedBigInt(newActivationDelay)
    )
  )

  return activationDelaySetEvent
}

export function createClaimerSetEvent(
  account: Address,
  claimer: Address
): ClaimerSet {
  let claimerSetEvent = changetype<ClaimerSet>(newMockEvent())

  claimerSetEvent.parameters = new Array()

  claimerSetEvent.parameters.push(
    new ethereum.EventParam("account", ethereum.Value.fromAddress(account))
  )
  claimerSetEvent.parameters.push(
    new ethereum.EventParam("claimer", ethereum.Value.fromAddress(claimer))
  )

  return claimerSetEvent
}

export function createGlobalCommissionBipsSetEvent(
  oldGlobalCommissionBips: i32,
  newGlobalCommissionBips: i32
): GlobalCommissionBipsSet {
  let globalCommissionBipsSetEvent = changetype<GlobalCommissionBipsSet>(
    newMockEvent()
  )

  globalCommissionBipsSetEvent.parameters = new Array()

  globalCommissionBipsSetEvent.parameters.push(
    new ethereum.EventParam(
      "oldGlobalCommissionBips",
      ethereum.Value.fromUnsignedBigInt(BigInt.fromI32(oldGlobalCommissionBips))
    )
  )
  globalCommissionBipsSetEvent.parameters.push(
    new ethereum.EventParam(
      "newGlobalCommissionBips",
      ethereum.Value.fromUnsignedBigInt(BigInt.fromI32(newGlobalCommissionBips))
    )
  )

  return globalCommissionBipsSetEvent
}

export function createPaymentClaimedEvent(
  token: Address,
  claimer: Address,
  amount: BigInt
): PaymentClaimed {
  let paymentClaimedEvent = changetype<PaymentClaimed>(newMockEvent())

  paymentClaimedEvent.parameters = new Array()

  paymentClaimedEvent.parameters.push(
    new ethereum.EventParam("token", ethereum.Value.fromAddress(token))
  )
  paymentClaimedEvent.parameters.push(
    new ethereum.EventParam("claimer", ethereum.Value.fromAddress(claimer))
  )
  paymentClaimedEvent.parameters.push(
    new ethereum.EventParam("amount", ethereum.Value.fromUnsignedBigInt(amount))
  )

  return paymentClaimedEvent
}

export function createPaymentUpdaterSetEvent(
  oldPaymentUpdater: Address,
  newPaymentUpdater: Address
): PaymentUpdaterSet {
  let paymentUpdaterSetEvent = changetype<PaymentUpdaterSet>(newMockEvent())

  paymentUpdaterSetEvent.parameters = new Array()

  paymentUpdaterSetEvent.parameters.push(
    new ethereum.EventParam(
      "oldPaymentUpdater",
      ethereum.Value.fromAddress(oldPaymentUpdater)
    )
  )
  paymentUpdaterSetEvent.parameters.push(
    new ethereum.EventParam(
      "newPaymentUpdater",
      ethereum.Value.fromAddress(newPaymentUpdater)
    )
  )

  return paymentUpdaterSetEvent
}

export function createRootSubmittedEvent(
  root: Bytes,
  paymentsCalculatedUntilTimestamp: BigInt,
  activatedAfter: BigInt
): RootSubmitted {
  let rootSubmittedEvent = changetype<RootSubmitted>(newMockEvent())

  rootSubmittedEvent.parameters = new Array()

  rootSubmittedEvent.parameters.push(
    new ethereum.EventParam("root", ethereum.Value.fromFixedBytes(root))
  )
  rootSubmittedEvent.parameters.push(
    new ethereum.EventParam(
      "paymentsCalculatedUntilTimestamp",
      ethereum.Value.fromUnsignedBigInt(paymentsCalculatedUntilTimestamp)
    )
  )
  rootSubmittedEvent.parameters.push(
    new ethereum.EventParam(
      "activatedAfter",
      ethereum.Value.fromUnsignedBigInt(activatedAfter)
    )
  )

  return rootSubmittedEvent
}
