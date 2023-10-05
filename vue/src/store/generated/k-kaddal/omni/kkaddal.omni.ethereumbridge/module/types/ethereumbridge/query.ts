/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../ethereumbridge/params";
import { EthState } from "../ethereumbridge/eth_state";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

export const protobufPackage = "kkaddal.omni.ethereumbridge";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetEthStateRequest {
  index: string;
}

export interface QueryGetEthStateResponse {
  ethState: EthState | undefined;
}

export interface QueryAllEthStateRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllEthStateResponse {
  ethState: EthState[];
  pagination: PageResponse | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetEthStateRequest: object = { index: "" };

export const QueryGetEthStateRequest = {
  encode(
    message: QueryGetEthStateRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetEthStateRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetEthStateRequest,
    } as QueryGetEthStateRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetEthStateRequest {
    const message = {
      ...baseQueryGetEthStateRequest,
    } as QueryGetEthStateRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetEthStateRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetEthStateRequest>
  ): QueryGetEthStateRequest {
    const message = {
      ...baseQueryGetEthStateRequest,
    } as QueryGetEthStateRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetEthStateResponse: object = {};

export const QueryGetEthStateResponse = {
  encode(
    message: QueryGetEthStateResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.ethState !== undefined) {
      EthState.encode(message.ethState, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetEthStateResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetEthStateResponse,
    } as QueryGetEthStateResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ethState = EthState.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetEthStateResponse {
    const message = {
      ...baseQueryGetEthStateResponse,
    } as QueryGetEthStateResponse;
    if (object.ethState !== undefined && object.ethState !== null) {
      message.ethState = EthState.fromJSON(object.ethState);
    } else {
      message.ethState = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetEthStateResponse): unknown {
    const obj: any = {};
    message.ethState !== undefined &&
      (obj.ethState = message.ethState
        ? EthState.toJSON(message.ethState)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetEthStateResponse>
  ): QueryGetEthStateResponse {
    const message = {
      ...baseQueryGetEthStateResponse,
    } as QueryGetEthStateResponse;
    if (object.ethState !== undefined && object.ethState !== null) {
      message.ethState = EthState.fromPartial(object.ethState);
    } else {
      message.ethState = undefined;
    }
    return message;
  },
};

const baseQueryAllEthStateRequest: object = {};

export const QueryAllEthStateRequest = {
  encode(
    message: QueryAllEthStateRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllEthStateRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllEthStateRequest,
    } as QueryAllEthStateRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllEthStateRequest {
    const message = {
      ...baseQueryAllEthStateRequest,
    } as QueryAllEthStateRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllEthStateRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllEthStateRequest>
  ): QueryAllEthStateRequest {
    const message = {
      ...baseQueryAllEthStateRequest,
    } as QueryAllEthStateRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllEthStateResponse: object = {};

export const QueryAllEthStateResponse = {
  encode(
    message: QueryAllEthStateResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.ethState) {
      EthState.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllEthStateResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllEthStateResponse,
    } as QueryAllEthStateResponse;
    message.ethState = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ethState.push(EthState.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllEthStateResponse {
    const message = {
      ...baseQueryAllEthStateResponse,
    } as QueryAllEthStateResponse;
    message.ethState = [];
    if (object.ethState !== undefined && object.ethState !== null) {
      for (const e of object.ethState) {
        message.ethState.push(EthState.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllEthStateResponse): unknown {
    const obj: any = {};
    if (message.ethState) {
      obj.ethState = message.ethState.map((e) =>
        e ? EthState.toJSON(e) : undefined
      );
    } else {
      obj.ethState = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllEthStateResponse>
  ): QueryAllEthStateResponse {
    const message = {
      ...baseQueryAllEthStateResponse,
    } as QueryAllEthStateResponse;
    message.ethState = [];
    if (object.ethState !== undefined && object.ethState !== null) {
      for (const e of object.ethState) {
        message.ethState.push(EthState.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a EthState by index. */
  EthState(request: QueryGetEthStateRequest): Promise<QueryGetEthStateResponse>;
  /** Queries a list of EthState items. */
  EthStateAll(
    request: QueryAllEthStateRequest
  ): Promise<QueryAllEthStateResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "kkaddal.omni.ethereumbridge.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  EthState(
    request: QueryGetEthStateRequest
  ): Promise<QueryGetEthStateResponse> {
    const data = QueryGetEthStateRequest.encode(request).finish();
    const promise = this.rpc.request(
      "kkaddal.omni.ethereumbridge.Query",
      "EthState",
      data
    );
    return promise.then((data) =>
      QueryGetEthStateResponse.decode(new Reader(data))
    );
  }

  EthStateAll(
    request: QueryAllEthStateRequest
  ): Promise<QueryAllEthStateResponse> {
    const data = QueryAllEthStateRequest.encode(request).finish();
    const promise = this.rpc.request(
      "kkaddal.omni.ethereumbridge.Query",
      "EthStateAll",
      data
    );
    return promise.then((data) =>
      QueryAllEthStateResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
