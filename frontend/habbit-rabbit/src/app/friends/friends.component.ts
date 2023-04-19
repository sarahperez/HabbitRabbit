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

  constructor (
    private friendService: FriendService
  ) { }

  onInit() {
    this.friendService
      .getFriendStatus();
    
  }

  requestsArr = sessionStorage.getItem('pendingRequests')

  onKey(event: any, type: string) {
    if (type === 'username') {
      this.friendusername = event.target.value;
    } 
  }

  onSubmit() {
    this.friendService
        .requestFriend(this.friendusername);
  }

}
