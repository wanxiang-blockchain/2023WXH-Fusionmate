import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { AccountService } from '../account/account.service';
import { BackendHost } from 'src/common/config';
import { AssistantType, BackendReturnType, CreateAIResponse, FetchAIListResponse, GenImgURIResponse, GetAIDetailResponse, HarvestResponse } from 'src/common/types';
import { BehaviorSubject, catchError, map, of } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class DataService {

  constructor(private http: HttpClient, private accountService: AccountService) { }

  // login(){
  // }

  genImgURI(prompt: string) {
    return this.http.post<GenImgURIResponse>(`${BackendHost}/api/v1/assistant/genImgURI`, { prompt }, {
      headers: { Token: this.accountService.getToken() }
    })
  }

  CreateAI(name: string, description: string, replication: number, mintPrice: string, prompt: string, imgURI: string, type: AssistantType, derive?: string) {
    return this.http.post<CreateAIResponse>(`${BackendHost}/api/v1/assistant/create`, {
      name, description, derive, replication, mintPrice, prompt, imgURI, type
    }, {
      headers: { Token: this.accountService.getToken() },
    })
  }

  notifyCreateResult(collectionID: string, contractAddress: string, result: boolean) {
    return this.http.post<BackendReturnType<null>>(`${BackendHost}/api/v1/assistant/notifyCreateResult`, {
      collectionID, result, contractAddress
    }, {
      headers: { Token: this.accountService.getToken() },
    })
  }

  fetchAIList(page: number, perPage: number, type?: string, makerAddr?: string) {
    const params: { [key: string]: any } = { page, perPage };
    if (type) {
      params['type'] = type;
    }
    if (makerAddr) {
      params['makerAddr'] = makerAddr;
    }
    return this.http.get<FetchAIListResponse>(`${BackendHost}/api/v1/assistant/collections`, {
      params,
      // headers: { token: this.accountService.getToken() }
    }).pipe(catchError(() => of({
      returnCode: 500,
      message: 'error',
      data: []
    } as FetchAIListResponse)))
  }

  getAIDetail(collectionID: string) {
    return this.http.get<GetAIDetailResponse>(`${BackendHost}/api/v1/assistant/collection/${collectionID}`, {
      // headers: { token: this.accountService.getToken() }
    })
  }

  sendMsg(collectionID: string, tokenId: string, msg: string) {
    const msgResponse = new BehaviorSubject('');

    fetch(`${BackendHost}/api/v1/chat/sendMsg`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Token': this.accountService.getToken()
      },
      body: JSON.stringify({
        collectionID,
        tokenId: Number(tokenId),
        msg
      })
    }).then(
      response => {
        if (response.status == 200) {
          return response.body?.getReader();
        }
        throw new Error(response.statusText)
      }
    ).then(
      async reader => {
        if (!reader) {
          return;
        }
        while (true) {
          const { done, value } = await reader.read();
          if (done) {
            msgResponse.complete();
            break;
          }
          let result = new TextDecoder().decode(value)
          console.log(('[' + result.replaceAll('}{', '},{') + ']'));
          const results = (JSON.parse(('[' + result.replaceAll('}{', '},{') + ']')) as { text: string }[] || []).map((item: { text: string }) => item.text).join('');
          msgResponse.next(msgResponse.value + results);
        }
      }
    ).catch(e => console.log(e));
    return msgResponse.pipe(map(msg => msg.trim()));

  }

  harvest(collectionID: string, tokenId: string) {
    return this.http.get<HarvestResponse>(`${BackendHost}/api/v1/harvest/${collectionID}/${tokenId}`, {
      headers: { Token: this.accountService.getToken() }
    })
  }
}
