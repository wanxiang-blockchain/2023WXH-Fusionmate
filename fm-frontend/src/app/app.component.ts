import { ChangeDetectorRef, Component, ElementRef, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { NavigationEnd, Router } from '@angular/router';
import { BehaviorSubject, Observable, Subscription, tap } from 'rxjs';
import { AccountService } from 'src/services/account/account.service';
import { WalletService } from 'src/services/wallet/wallet.service';
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit, OnDestroy {
  // currentAccount$: BehaviorSubject<string | undefined>;

  currentAccount: string | undefined;
  subCurrentAccount?: Subscription;
  subRouter?: Subscription;
  @ViewChild('routeContainer')
  routeContainerRef?: ElementRef;
  constructor(private walletService: WalletService, private accountService: AccountService, private changeDetectorRef: ChangeDetectorRef, private router: Router) {
  }
  ngOnDestroy(): void {
    this.subCurrentAccount?.unsubscribe();
    this.subRouter?.unsubscribe();
  }
  ngOnInit(): void {
    // this.connectWallet();
    this.subCurrentAccount = this.accountService.account$.subscribe(
      account => {
        this.currentAccount = account ? this.formatAddress(account) : undefined;
        this.changeDetectorRef.detectChanges();
      }
    );

    this.subRouter = this.router.events.subscribe((event) => {
      if (event instanceof NavigationEnd) {
        this.routeContainerRef?.nativeElement.scrollTo(0, 0);
      }
    });
  }

  connectWallet() {
    this.walletService.connect();
  }
  formatAddress(address: string) {
    return address.substring(0, 10) + '...' + address.substring(32)
  }
}
