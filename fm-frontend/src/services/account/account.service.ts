import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AccountService {


  public account$ = new BehaviorSubject<string | undefined>(undefined);
  private token = '';
  constructor() { }
  setToken(token: string) {
    this.token = token;
    sessionStorage.setItem(`${this.account$.value}_ctoken`, token);
  }
  getToken(address?: string) {
    this.token = sessionStorage.getItem(`${address?.toLowerCase() || this.account$.value}_ctoken`)||'' ;
    return this.token;
  }

}
