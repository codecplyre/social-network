import './App.css';
import Footer from './layouts/Footer/Footer';
import Header from './layouts/Header/Header';
import Home from './pages/Home';
import { Route, Routes } from 'react-router-dom';
import Login from './pages/loginPage/Login';
import Registration from './pages/regPage/Registration';
import Chat from './features/Chat/Chat';
import Profile from './pages/profile/Profile';
import Stories from './pages/stories/stories';
import Comments from './features/Comments';
import { useEffect, useState } from 'react';
import NewPost from './features/newpost/NewPost';
import { Notification } from './features/Notification/Notification';
import { Search } from './features/Search/Search';
import Validation from './components/Validation/Validation';
function App() {
    const [isMobile, setIsMobile] = useState(false);
    const [authorised, setAuthorised] = useState(false);
    let socket;
    //random userid
    const user = Math.floor(Math.random() * 5);
    let generalNotif = [
        {
            name: 'John',
            id: 1,
            content: 'liked your post',
            optContent: '1h',
        },
    ];
    let followrequest = [{ name: 'Ken' }];
    let groupNotif = [];

    // for websocket
    const CreateWebSocket = () => {
        socket = new WebSocket('ws://' + 'localhost:5070' + '/ws');
        console.log('Attempting Connection...');
    };
    CreateWebSocket();
    return (
        <div
            className='App'
            ref={(boxRef) => {
                boxRef &&
                    console.log(
                        boxRef.getBoundingClientRect().width,
                        boxRef.getBoundingClientRect().width >= 600
                    );
                return (
                    boxRef &&
                    setIsMobile(boxRef.getBoundingClientRect().width < 600)
                );
            }}>
            {
                <Validation>
                    <Header />
                </Validation>
            }

            <>
                <Routes>
                    <Route path='/' element={<Login auth={setAuthorised} />} />
                    <Route path='/register' element={<Registration />} />

                    <Route
                        path='/home'
                        element={
                            isMobile ? (
                                <Validation>
                                    <Home
                                        bodyStyleName='mobile'
                                        cardStyleName='mobileCard'
                                    />
                                </Validation>
                            ) : (
                                <Validation>
                                    <Home
                                        bodyStyleName='desktop'
                                        cardStyleName='desktopCard'
                                    />
                                </Validation>
                            )
                        }
                    />
                    <Route
                        path='/search'
                        element={
                            <Validation>
                                <Search />
                            </Validation>
                        }
                    />
                    <Route
                        path='/newpost'
                        element={
                            <Validation>
                                <NewPost />
                            </Validation>
                        }
                    />
                    <Route
                        path='/messages'
                        element={
                            <Validation>
                                isMobile ? (
                                <Chat
                                    bodyStyleName='mobile'
                                    cardStyleName='mobileCard'
                                    user={user}
                                    socket={socket}
                                />
                                ) : (
                                <Chat
                                    bodyStyleName='desktop'
                                    cardStyleName='desktopCard'
                                    user={user}
                                    socket={socket}
                                />
                                )
                            </Validation>
                        }
                    />
                    <Route
                        path='/groups'
                        element={
                            <Validation>
                                <h1>Groups</h1>
                            </Validation>
                        }
                    />
                    <Route
                        path='/comments'
                        element={
                            <Validation>
                                isMobile ? (
                                <Comments
                                    bodyStyleName='mobile'
                                    cardStyleName='mobileCard'
                                />
                                ) : (
                                <Comments
                                    bodyStyleName='desktop'
                                    cardStyleName='desktopCard'
                                />
                                )
                            </Validation>
                        }
                    />
                    <Route
                        path='notifications'
                        element={
                            <Validation>
                                <Notification users={generalNotif} />
                            </Validation>
                        }>
                        <Route
                            path='general'
                            element={
                                <Validation>
                                    <Notification users={generalNotif} />
                                </Validation>
                            }
                        />
                        <Route
                            path='followrequest'
                            element={
                                <Validation>
                                    <Notification users={followrequest} />
                                </Validation>
                            }
                        />
                        <Route
                            path='group'
                            element={
                                <Validation>
                                    <Notification users={groupNotif} />
                                </Validation>
                            }
                        />
                    </Route>
                    <Route
                        path='profile'
                        element={
                            <Validation>
                                <Profile
                                    aboutMe='This section is where the bio goes. You should write 1-2 sentences about yourself.'
                                    fullname='Nathaniel Russell'
                                    nickname='Nate'
                                />
                            </Validation>
                        }
                    />
                    <Route
                        path='/stories'
                        element={
                            <Validation>
                                <Stories />
                            </Validation>
                        }
                    />
                </Routes>
            </>
            {
                <Validation>
                    <Footer />
                </Validation>
            }
        </div>
    );
}

export default App;
