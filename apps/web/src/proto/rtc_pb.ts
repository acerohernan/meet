// @generated by protoc-gen-es v1.6.0 with parameter "target=ts"
// @generated from file rtc.proto (package core, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message core.SignalRequest
 */
export class SignalRequest extends Message<SignalRequest> {
  /**
   * @generated from oneof core.SignalRequest.request
   */
  request: {
    /**
     * @generated from field: core.StartSession start_session = 1;
     */
    value: StartSession;
    case: "startSession";
  } | { case: undefined; value?: undefined } = { case: undefined };

  /**
   * @generated from field: string participant_id = 2;
   */
  participantId = "";

  constructor(data?: PartialMessage<SignalRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "core.SignalRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "start_session", kind: "message", T: StartSession, oneof: "request" },
    { no: 2, name: "participant_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SignalRequest {
    return new SignalRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SignalRequest {
    return new SignalRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SignalRequest {
    return new SignalRequest().fromJsonString(jsonString, options);
  }

  static equals(a: SignalRequest | PlainMessage<SignalRequest> | undefined, b: SignalRequest | PlainMessage<SignalRequest> | undefined): boolean {
    return proto3.util.equals(SignalRequest, a, b);
  }
}

/**
 * @generated from message core.SignalResponse
 */
export class SignalResponse extends Message<SignalResponse> {
  /**
   * @generated from oneof core.SignalResponse.response
   */
  response: {
    /**
     * @generated from field: core.JoinResponse join_response = 1;
     */
    value: JoinResponse;
    case: "joinResponse";
  } | { case: undefined; value?: undefined } = { case: undefined };

  /**
   * @generated from field: string participant_id = 2;
   */
  participantId = "";

  constructor(data?: PartialMessage<SignalResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "core.SignalResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "join_response", kind: "message", T: JoinResponse, oneof: "response" },
    { no: 2, name: "participant_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SignalResponse {
    return new SignalResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SignalResponse {
    return new SignalResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SignalResponse {
    return new SignalResponse().fromJsonString(jsonString, options);
  }

  static equals(a: SignalResponse | PlainMessage<SignalResponse> | undefined, b: SignalResponse | PlainMessage<SignalResponse> | undefined): boolean {
    return proto3.util.equals(SignalResponse, a, b);
  }
}

/**
 * @generated from message core.StartSession
 */
export class StartSession extends Message<StartSession> {
  constructor(data?: PartialMessage<StartSession>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "core.StartSession";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StartSession {
    return new StartSession().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StartSession {
    return new StartSession().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StartSession {
    return new StartSession().fromJsonString(jsonString, options);
  }

  static equals(a: StartSession | PlainMessage<StartSession> | undefined, b: StartSession | PlainMessage<StartSession> | undefined): boolean {
    return proto3.util.equals(StartSession, a, b);
  }
}

/**
 * @generated from message core.JoinResponse
 */
export class JoinResponse extends Message<JoinResponse> {
  constructor(data?: PartialMessage<JoinResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "core.JoinResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): JoinResponse {
    return new JoinResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): JoinResponse {
    return new JoinResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): JoinResponse {
    return new JoinResponse().fromJsonString(jsonString, options);
  }

  static equals(a: JoinResponse | PlainMessage<JoinResponse> | undefined, b: JoinResponse | PlainMessage<JoinResponse> | undefined): boolean {
    return proto3.util.equals(JoinResponse, a, b);
  }
}

