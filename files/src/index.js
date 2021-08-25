import React from "react";
import ReactDOM from "react-dom";
import "./sass/style.scss";

function Background() {
    return (
        <span>
            <span className="bg"></span>
            <span className="bg bg2"></span>
            <span className="bg bg3"></span>
        </span>
    );
}

function Product(props) {
    return (
        <li className="products__item">
            <ul className="products__attributes">
                <li id="img-btn" className="products__image" data-id-product={props.id}>
                    <img className="products__img" alt="xd" src="https://picsum.photos/200/200" />
                    <span className="products__img-label">Buy Now</span>
                </li>
                <li id="name-btn" className="products__name" data-id-product={props.id}>{props.name}</li>
                <li id="price-btn" className="products__price" data-id-product={props.id}>PRICE: {props.price}$</li>
                <li className="products__description"><h4>Description</h4>{props.type} {props.brand} Price:{props.price}$ Stock: {props.stock}</li>
            </ul>
        </li>
    )
}

class ListOfProducts extends React.Component {
    render() {
        return (
            <ul className="products">
                {this.props.products.map((product, index) => {
                    return <Product id={product.id} name={product.name} price={product.price} type={product.type} brand={product.brand} stock={product.stock} />
                })}
            </ul>
        )
    }
}

class Form extends React.Component {
    constructor(props) {
        super(props);
        this.state = { value: "" };

        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleChange(event) {
        this.setState({ value: event.target.value });
    }

    handleSubmit(event) {
        event.preventDefault();
        const search = { sentence: this.state.value };

        fetch("/api/v1/search", {
            method: "POST",
            body: JSON.stringify(search),
        })
            .then((responsive) => {
                if (responsive.status >= 400) {
                    throw true;
                }
                return responsive.json();
            })
            .then((products) => {
                ReactDOM.render(<div>
                    <Background />
                    <Index />
                    <ListOfProducts products={products} />
                </div >, document.getElementById("root"));
            })
            .catch(() => {
                ReactDOM.render(<div>
                    <Background />
                    <Index />
                    <p>Error Not Found</p>
                </div >, document.getElementById("root"));
            });
    }

    render() {
        return (
            <form onSubmit={this.handleSubmit}>
                <input
                    type="text"
                    value={this.state.value}
                    onChange={this.handleChange}
                    required
                />
                <input type="submit" value="Search" />
            </form>
        );
    }
}

class Index extends React.Component {
    render() {
        return (
            <div>
                <Background />
                <h1 className="title">Search Engine Test</h1>
                <Form />
            </div>
        );
    }
}

ReactDOM.render(<Index />, document.getElementById("root"));
