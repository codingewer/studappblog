import './App.css';
import { Routes, Route, useParams } from "react-router-dom";

import Home from './layouts/Home';
import ButtomNav from './bars/ButtomNav';
import { BrowserRouter as Router } from "react-router-dom";
import ExplorePage from './layouts/ExplorePage';
import Navi from './bars/Navi';
import Signup from './User/Register';
import Login from './User/Login';
import Edit from './posts/Edit';
import AdminPanel from './admin/AdminPanel';
import Profile from './User/Profile';
import EditUser from "./User/EditUser"
import UploadImage from './image/UploadImage';
import UpdateUserAvatar from './User/UpdateUserAvatar';
import UploadPostImg from './posts/UploadPostImg';
import UpdatePost from './posts/UpdatePost';
import DeletePost from './posts/DeletePost';
import UserPosts from "./posts/UserPosts"
import Post from './posts/Post';


function App() {
  return (
    < div  >
      <Router>
        <Navi />
        <ButtomNav />
        <Routes>
          <Route path='/' element={<Home />} />
          <Route path='/home' element={<Home />} />
          <Route path='/explore' element={<ExplorePage />} />
          <Route path='/register' element={<Signup />} />
          <Route path='/login' element={<Login />} />
          <Route path='/editor' element={<Edit />} />
          <Route path='/adminpanel' element={<AdminPanel />} />
          <Route path='/profile' element={<Profile />} />
          <Route path='/editUser' element={<EditUser />} />
          <Route path='/uploadImage' element={<UploadImage />} />
          <Route path='/useravatar' element={<UpdateUserAvatar />} />
          <Route path='/postimg' element={<UploadPostImg />} />
          <Route path='/userPosts/:userId' element={<UserPosts />} />
          <Route path='/updatePost/:postId' element={<UpdatePost />} />
          <Route path='/deletepost/:postId' element={<DeletePost />} />
          <Route path='/post/:postId' element={<Post />} />
        </Routes>
        <ButtomNav />
      </Router>
    </div>
  );
}

export default App;
