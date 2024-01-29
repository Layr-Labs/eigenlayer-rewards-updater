import {
  assert,
  describe,
  test,
  clearStore,
  beforeAll,
  afterAll
} from "matchstick-as/assembly/index"
import {} from "@graphprotocol/graph-ts"
import { RangePaymentCreated } from "../generated/schema"
import { RangePaymentCreated as RangePaymentCreatedEvent } from "../generated/PaymentCoordinator/PaymentCoordinator"
import { handleRangePaymentCreated } from "../src/payment-coordinator"
import { createRangePaymentCreatedEvent } from "./payment-coordinator-utils"

// Tests structure (matchstick-as >=0.5.0)
// https://thegraph.com/docs/en/developer/matchstick/#tests-structure-0-5-0

describe("Describe entity assertions", () => {
  beforeAll(() => {
    let rangePayment = "ethereum.Tuple Not implemented"
    let newRangePaymentCreatedEvent =
      createRangePaymentCreatedEvent(rangePayment)
    handleRangePaymentCreated(newRangePaymentCreatedEvent)
  })

  afterAll(() => {
    clearStore()
  })

  // For more test scenarios, see:
  // https://thegraph.com/docs/en/developer/matchstick/#write-a-unit-test

  test("RangePaymentCreated created and stored", () => {
    assert.entityCount("RangePaymentCreated", 1)

    // 0xa16081f360e3847006db660bae1c6d1b2e17ec2a is the default address used in newMockEvent() function
    assert.fieldEquals(
      "RangePaymentCreated",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "rangePayment",
      "ethereum.Tuple Not implemented"
    )

    // More assert options:
    // https://thegraph.com/docs/en/developer/matchstick/#asserts
  })
})
