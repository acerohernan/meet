// @generated by protoc-gen-es v1.6.0 with parameter "target=ts"
// @generated from file guest.proto (package core, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message core.GuestJoinRequest
 */
export class GuestJoinRequest extends Message<GuestJoinRequest> {
  /**
   * @generated from field: string room_id = 1;
   */
  roomId = "";

  /**
   * @generated from field: core.Guest guest = 2;
   */
  guest?: Guest;

  constructor(data?: PartialMessage<GuestJoinRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "core.GuestJoinRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "room_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "guest", kind: "message", T: Guest },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GuestJoinRequest {
    return new GuestJoinRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GuestJoinRequest {
    return new GuestJoinRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GuestJoinRequest {
    return new GuestJoinRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GuestJoinRequest | PlainMessage<GuestJoinRequest> | undefined, b: GuestJoinRequest | PlainMessage<GuestJoinRequest> | undefined): boolean {
    return proto3.util.equals(GuestJoinRequest, a, b);
  }
}

/**
 * @generated from message core.GuestRequestCancelled
 */
export class GuestRequestCancelled extends Message<GuestRequestCancelled> {
  /**
   * @generated from field: string room_id = 1;
   */
  roomId = "";

  /**
   * @generated from field: string guest_id = 2;
   */
  guestId = "";

  constructor(data?: PartialMessage<GuestRequestCancelled>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "core.GuestRequestCancelled";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "room_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "guest_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GuestRequestCancelled {
    return new GuestRequestCancelled().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GuestRequestCancelled {
    return new GuestRequestCancelled().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GuestRequestCancelled {
    return new GuestRequestCancelled().fromJsonString(jsonString, options);
  }

  static equals(a: GuestRequestCancelled | PlainMessage<GuestRequestCancelled> | undefined, b: GuestRequestCancelled | PlainMessage<GuestRequestCancelled> | undefined): boolean {
    return proto3.util.equals(GuestRequestCancelled, a, b);
  }
}

/**
 * @generated from message core.GuestJoinResponse
 */
export class GuestJoinResponse extends Message<GuestJoinResponse> {
  /**
   * @generated from field: core.Guest guest = 1;
   */
  guest?: Guest;

  /**
   * @generated from oneof core.GuestJoinResponse.answer
   */
  answer: {
    /**
     * @generated from field: core.JoinApproved join_approved = 2;
     */
    value: JoinApproved;
    case: "joinApproved";
  } | {
    /**
     * @generated from field: core.JoinDenied join_denied = 3;
     */
    value: JoinDenied;
    case: "joinDenied";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<GuestJoinResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "core.GuestJoinResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "guest", kind: "message", T: Guest },
    { no: 2, name: "join_approved", kind: "message", T: JoinApproved, oneof: "answer" },
    { no: 3, name: "join_denied", kind: "message", T: JoinDenied, oneof: "answer" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GuestJoinResponse {
    return new GuestJoinResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GuestJoinResponse {
    return new GuestJoinResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GuestJoinResponse {
    return new GuestJoinResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GuestJoinResponse | PlainMessage<GuestJoinResponse> | undefined, b: GuestJoinResponse | PlainMessage<GuestJoinResponse> | undefined): boolean {
    return proto3.util.equals(GuestJoinResponse, a, b);
  }
}

/**
 * @generated from message core.JoinApproved
 */
export class JoinApproved extends Message<JoinApproved> {
  /**
   * @generated from field: string access_token = 1;
   */
  accessToken = "";

  constructor(data?: PartialMessage<JoinApproved>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "core.JoinApproved";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "access_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): JoinApproved {
    return new JoinApproved().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): JoinApproved {
    return new JoinApproved().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): JoinApproved {
    return new JoinApproved().fromJsonString(jsonString, options);
  }

  static equals(a: JoinApproved | PlainMessage<JoinApproved> | undefined, b: JoinApproved | PlainMessage<JoinApproved> | undefined): boolean {
    return proto3.util.equals(JoinApproved, a, b);
  }
}

/**
 * @generated from message core.JoinDenied
 */
export class JoinDenied extends Message<JoinDenied> {
  constructor(data?: PartialMessage<JoinDenied>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "core.JoinDenied";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): JoinDenied {
    return new JoinDenied().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): JoinDenied {
    return new JoinDenied().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): JoinDenied {
    return new JoinDenied().fromJsonString(jsonString, options);
  }

  static equals(a: JoinDenied | PlainMessage<JoinDenied> | undefined, b: JoinDenied | PlainMessage<JoinDenied> | undefined): boolean {
    return proto3.util.equals(JoinDenied, a, b);
  }
}

/**
 * @generated from message core.Guest
 */
export class Guest extends Message<Guest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * @generated from field: string node_id = 3;
   */
  nodeId = "";

  constructor(data?: PartialMessage<Guest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "core.Guest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "node_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Guest {
    return new Guest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Guest {
    return new Guest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Guest {
    return new Guest().fromJsonString(jsonString, options);
  }

  static equals(a: Guest | PlainMessage<Guest> | undefined, b: Guest | PlainMessage<Guest> | undefined): boolean {
    return proto3.util.equals(Guest, a, b);
  }
}

