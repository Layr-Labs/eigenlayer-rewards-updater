import { newMockEvent } from "matchstick-as"
import { ethereum } from "@graphprotocol/graph-ts"
import { RangePaymentCreated } from "../generated/PaymentCoordinator/PaymentCoordinator"

export function createRangePaymentCreatedEvent(
  rangePayment: ethereum.Tuple
): RangePaymentCreated {
  let rangePaymentCreatedEvent = changetype<RangePaymentCreated>(newMockEvent())

  rangePaymentCreatedEvent.parameters = new Array()

  rangePaymentCreatedEvent.parameters.push(
    new ethereum.EventParam(
      "rangePayment",
      ethereum.Value.fromTuple(rangePayment)
    )
  )

  return rangePaymentCreatedEvent
}
