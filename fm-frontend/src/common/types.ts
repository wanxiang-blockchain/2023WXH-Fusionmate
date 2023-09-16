import type { Observable } from "rxjs";
import type { TransactionReceipt, AbiParameter } from "web3-types";

export type ContractSendType = 'send' | 'call';

export type ContractReturnType<T> = T extends 'send' ? TransactionReceipt : T extends 'call' ? { [key: string]: any; __length__: number; } : never;

export type BackendReturnType<T> = {
    returnCode: number,
    message: string,
    data: T
};

export type LoginResponse = BackendReturnType<{
    token: string
}>;

export type GenImgURIResponse = BackendReturnType<{
    imgURI: string
}>;

export type CreateAIResponse = BackendReturnType<{
    collectionID: string
    name: string
    symbol: string
    baseURI: string
    maxSupply: string
    mintPrice: string
    signature: string
}>;

export type FetchAIListResponse = BackendReturnType<{
    collectionID: string
    name: string
    symbol: string
    baseURI: string
    maxSupply: string
    mintPrice: string
    description: string
    derive?: string
    prompts: string
    contractAddr: string
    maker: string
    imgURI: string
    type: AssistantType
}[]>;

export type GetAIDetailResponse = BackendReturnType<{
    collectionID: string
    name: string
    symbol: string
    baseURI: string
    maxSupply: string
    mintPrice: string
    description: string
    derive?: string
    prompts: string
    contractAddr: string
    maker: string
    imgURI: string
    type: AssistantType
}>;

export type GetMetadataResponse = BackendReturnType<{
    description: string
    externalURL?: string
    image: string
    name: string
    attributes: {
        traitTypeP: string
        value: string
    }[]
}>;

export type SendMsgResponse = BackendReturnType<{
    msg: string
}>;

export type HarvestResponse = BackendReturnType<{
    collectionID: string,
    tokenID: string,
    tokenNum: string
    signature: string
}>;

export type Assistant = {
    collectionID: string
    name: string
    symbol: string
    baseURI: string
    maxSupply: string
    mintPrice: string
    description: string
    derive?: string
    prompts: string
    contractAddr: string
    maker: string
    imgURI: string
    type: AssistantType
    mintAmount$: Observable<number>
    owned$: Observable<boolean>
    isMaker$: Observable<boolean>
    ownedTokenId$: Observable<string[]>
};

// 0:web3 expert  1:game characters 2:novel character 3:celebrities 4:Anime girl
export type AssistantType = '0' | '1' | '2' | '3' | '4';