import { Component } from "react";
import Link from "next/link";
import fetch from "isomorphic-unfetch";
import axios from "axios";
import Header from "../components/Header";

class ChatOne extends Component {
  // init state with the prefetched messages
  constructor(props) {
    super(props);
    this.state = {
      field: "",
      newMessage: 0,
      messages: [],
      subscribe: false,
      subscribed: false
    };
  }

  subscribe = () => {
    if (this.state.subscribe && !this.state.subscribed) {
      // connect to WS server and listen event
      this.props.socket.on("message.chat1", this.handleMessage);
      this.setState({ subscribed: true });
    }
  };

  componentDidMount() {
    this.subscribe();
    const token = window.localStorage.getItem("token");

    const data = {
      headers: {
        Authorization: `Bearer ${token}`
      },
      params: { room_id: "0" }
    };

    (async () => {
      await axios.get("http://localhost:3001/messages", data).then(response => {
        this.setState({ messages: response.data });
      });
    })();
  }

  componentDidUpdate() {
    this.subscribe();
  }

  static getDerivedStateFromProps(props, state) {
    if (props.socket && !state.subscribe) return { subscribe: true };
    return null;
  }

  // close socket connection
  componentWillUnmount() {
    this.props.socket.off("message.chat1", this.handleMessage);
  }

  // add messages from server to the state
  handleMessage = message => {
    this.setState(state => ({ messages: state.messages.concat(message) }));
  };

  handleOtherMessage = () => {
    this.setState(prevState => ({ newMessage: prevState.newMessage + 1 }));
  };

  handleChange = event => {
    this.setState({ field: event.target.value });
  };

  // send messages to server and add them to the state
  handleSubmit = event => {
    event.preventDefault();

    const token = window.localStorage.getItem("token");

    // create message object
    const message = {
      headers: {
        Authorization: `Bearer ${token}`
      },
      id: new Date().getTime(),
      content: this.state.field
    };

    // send object to WS server
    this.props.socket.emit("message.chat1", message);
    axios.post("http://localhost:3001/messages", message);

    // add it to state and clean current input value
    this.setState(state => ({
      field: "",
      messages: state.messages.concat(message)
    }));
  };

  render() {
    return (
      <>
        <Header>Ask Me Anything</Header>
        <main>
          <div>
            <Link href="/home">
              <a>他のルームへ</a>
            </Link>
            <ul>
              {this.state.messages.map((message, index) => (
                <li key={index}>
                  <div>ユーザー名が入ります</div>
                  {message.content}
                </li>
              ))}
            </ul>
            <form onSubmit={e => this.handleSubmit(e)}>
              <input
                onChange={this.handleChange}
                type="text"
                placeholder="メッセージを入力"
                value={this.state.field}
              />
              <button>投稿</button>
            </form>
          </div>
        </main>
      </>
    );
  }
}

export default ChatOne;
