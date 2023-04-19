import { Component } from '@angular/core';
import { FriendService } from '../services/friend/friend.service';

interface SideNavToggle{
  screenWidth: number;
  collapsed: boolean;
}

@Component({
  selector: 'app-friends',
  templateUrl: './friends.component.html',
  styleUrls: ['./friends.component.css']
})

export class FriendsComponent {
  //Sidenav stuff
  isSideNavCollapsed = false;
  screenWidth=0;
  onToggleSideNav(data: SideNavToggle): void{
    this.screenWidth = data.screenWidth;
    this.isSideNavCollapsed = data.collapsed;
  }

  friendusername: any;
  pendingRequests: any;
  blocked: any;
  friends: any;

  constructor (
    private friendService: FriendService
  ) { }

  ngOnInit(): void {
    this.friendService
      .getFriendStatus();
    this.pendingRequests = JSON.parse(localStorage.getItem('pendingRequests')!);
    this.blocked = JSON.parse(localStorage.getItem('blocked')!);
    this.friends = JSON.parse(localStorage.getItem('friends')!);
  }

  onKey(event: any, type: string) {
    if (type === 'username') {
      this.friendusername = event.target.value;
    } 
  }

  acceptRequest(sender: string){
    this.friendService
      .acceptFriend(sender);

    this.friendService
      .getFriendStatus();

    this.pendingRequests = localStorage.getItem('pendingRequests');
    this.blocked = JSON.parse(localStorage.getItem('blocked')!);
    this.friends = JSON.parse(localStorage.getItem('friends')!);
  }

  blockRequest(sender: string){
    this.friendService
      .blockFriend(sender);

    this.friendService
      .getFriendStatus(); 

    this.pendingRequests = localStorage.getItem('pendingRequests');
    this.blocked = JSON.parse(localStorage.getItem('blocked')!);
    this.friends = JSON.parse(localStorage.getItem('friends')!);
  }

  onSubmit() {
    this.friendService
        .requestFriend(this.friendusername);
  }

}
