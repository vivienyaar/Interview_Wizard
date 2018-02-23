import React, {Component} from 'react';

class App extends Component {
	constructor() {
	    super();
	    this.ws = new WebSocket('ws://127.0.0.1:8080/');
	    this.state = {
	        tmpMsg: ""
	    };
    }

    componentDidMount() {
		this.ws.onopen = (e) => {
			console.log("socket open")
		};

        this.ws.onmessage = (e) => {
            this.setState({ tmpMsg: e.data }) ;
        };
    }

    componentWillUnmount() {
        if (this.ws) {
            this.ws.close();
        }
    }


	render() {
		return (
			<div>
				<h1>Hello Wizard</h1>
				<div>{this.state.tmpMsg}</div>
			</div>
		);
	}
}
export default App;
