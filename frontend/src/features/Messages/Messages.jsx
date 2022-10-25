import './Messages.css';
import SendMessageBox from './components/SendMessageBox';
import Input from '../../components/Input/Input';
import { MessagesIcon, UserIcon } from '../../components/Icons/Icons';
import { MessageContainer } from './components/messageContainer';
import { MessageContent } from './components/MessageContent';
import { ProfileIcon } from '../../components/Icons/Icons';
import { useRef, useState } from 'react';
import { useEffect } from 'react';

export const Messages = ({ name, img, msg, socket }) => {
    let messageInput = useRef();
    let chatBox = useRef();
    const [messages, setMessages] = useState(msg);
    const sendMessage = (e) => {
        e.preventDefault();
        console.log('hello e', e.target);
        if (messageInput.current.value !== '') {
            // console.log(
            //     JSON.stringify({
            //         type: 'message', // message, notification, followrequest
            //         receiver: name, //name of the receiver
            //         sender: 'Moment', // change this to current user
            //         img: img, // img of the sender
            //         content: messageInput.current.value, // content of the message
            //     })
            // );
            socket.send(
                JSON.stringify({
                    type: 'message', // message, notification, followrequest
                    receiver: name, //name of the receiver
                    sender: 'Moment', // change this to current user
                    img: img, // img of the sender
                    content: messageInput.current.value, // content of the message
                })
            );
            messageInput.current.value = '';
            chatBox.current.scroll({
                top: chatBox.current.scrollHeight,
            });
        }
    };
    useEffect(() => {
        chatBox.current.scroll({ top: chatBox.current.scrollHeight });
    }, [messages]);
    const handleKeyDown = (e) => {
        if (e.key === 'Enter') {
            sendMessage(e);
            return;
        }
        // socket.send(
        //     JSON.stringify({
        //         type: 'typing', // message, notification, followrequest
        //         body: "I'm typing",
        //     })
        // );
    };
    socket.onmessage = (event) => {
        setMessages((messages) => [...messages, JSON.parse(event.data)]);
    };
    return (
        <>
            <div className='chatMessageContainerHeader'>
                <div className='messageContainerHeaderContent'>
                    <div className='messageHeaderInfo'>
                        <span className='chat'>
                            {
                                <ProfileIcon
                                    iconStyleName='imgIcon'
                                    imgStyleName='imgIcon'
                                    img={img}
                                />
                            }
                        </span>
                        <span className='messageHeaderName longTextElipsis'>
                            {name}
                        </span>
                    </div>
                    {/* this will be replace by the elipsis btn */}
                    <MessagesIcon />
                </div>
            </div>
            <div className='chatMessageContainer' ref={chatBox}>
                {messages.map((message, i) => {
                    return (
                        <MessageContainer
                            key={message + i}
                            message={message}
                            name={name}>
                            <MessageContent
                                date={'11 October 2022 •17:46'}
                                content='Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed euismod, nisl sit amet aliquam aliquam, nisl nisl aliquam nisl, sit amet aliquam nisl nisl sit amet.'
                            />
                        </MessageContainer>
                    );
                })}
                {/* <MessageContainer
                    name='Firstname ajsdkjgasjdglajshgdjla;ksdklfnaskd;akldhk;asdk;asjgsjldgaljsgdljagsjldgasjldgajls'
                    img='./logo.svg'>
                    <MessageContent
                        date={'11 October 2022 •17:46'}
                        content='Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed euismod, nisl sit amet aliquam aliquam, nisl nisl aliquam nisl, sit amet aliquam nisl nisl sit amet.'
                    />
                </MessageContainer> */}
            </div>
            <div className='messageInputContainer'>
                {/* this will be replace by the emoji btn */}
                <div className='inputContainer'>
                    <div>
                        <MessagesIcon />
                    </div>
                    {/* <
                        placeholder='message'
                        styleName='messageInput'></Input> */}
                    <div className='messageInput'>
                        <textarea
                            type='submit'
                            id=''
                            rows='2'
                            placeholder='message'
                            ref={messageInput}
                            onKeyDown={handleKeyDown}></textarea>
                    </div>
                    <div onClick={sendMessage}>
                        <MessagesIcon />
                    </div>
                </div>
            </div>
        </>
    );
};
