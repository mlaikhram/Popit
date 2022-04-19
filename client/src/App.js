import { Button, Card, CardBody, CardTitle, Col, Input, InputGroup, Row, Spinner, UncontrolledCollapse } from "reactstrap"
import { PageCard } from "./components/PageCard";
import { PageScrollMenu } from "./components/PageScrollMenu";
import { Link } from "react-router-dom";

import logo from './logo.svg';
import './App.css';


const items = [
  {
    "id": "625798abc15bd223b8a9eeae",
    "name": "Attack on Titan",
    "aliases": [
        "Shingeki no Kyojin"
    ],
    "images": {
        "1": "https://cdn.myanimelist.net/images/anime/5/44560l.jpg"
    }
  },
  {
    "id": "625798abc15bd223b8a9eeae1",
    "name": "Attack on Titan",
    "aliases": [
        "Shingeki no Kyojin"
    ],
    "images": {
        "1": "https://cdn.myanimelist.net/images/anime/5/44560l.jpg"
    }
  },
  {
    "id": "625798abc15bd223b8a9eeae2",
    "name": "Attack on Titan",
    "aliases": [
        "Shingeki no Kyojin"
    ],
    "images": {
        "1": "https://cdn.myanimelist.net/images/anime/5/44560l.jpg"
    }
  },
  {
    "id": "625798abc15bd223b8a9eeae3",
    "name": "Attack on Titan",
    "aliases": [
        "Shingeki no Kyojin"
    ],
    "images": {
        "1": "https://cdn.myanimelist.net/images/anime/5/44560l.jpg"
    }
  },
  {
    "id": "625798abc15bd223b8a9eeae4",
    "name": "Attack on Titan",
    "aliases": [
        "Shingeki no Kyojin"
    ],
    "images": {
        "1": "https://cdn.myanimelist.net/images/anime/5/44560l.jpg"
    }
  },
  {
    "id": "625798abc15bd223b8a9eeae5",
    "name": "Attack on Titan",
    "aliases": [
        "Shingeki no Kyojin"
    ],
    "images": {
        "1": "https://cdn.myanimelist.net/images/anime/5/44560l.jpg"
    }
  },
  {
    "id": "625798abc15bd223b8a9eeae6",
    "name": "Attack on Titan",
    "aliases": [
        "Shingeki no Kyojin"
    ],
    "images": {
        "1": "https://cdn.myanimelist.net/images/anime/5/44560l.jpg"
    }
  },
  {
    "id": "625798abc15bd223b8a9eeae7",
    "name": "Attack on Titan",
    "aliases": [
        "Shingeki no Kyojin"
    ],
    "images": {
        "1": "https://cdn.myanimelist.net/images/anime/5/44560l.jpg"
    }
  },
  {
    "id": "625798abc15bd223b8a9eeae8",
    "name": "Attack on Titan",
    "aliases": [
        "Shingeki no Kyojin"
    ],
    "images": {
        "1": "https://cdn.myanimelist.net/images/anime/5/44560l.jpg"
    }
  },
  {
    "id": "625798abc15bd223b8a9eeae9",
    "name": "Attack on Titan",
    "aliases": [
        "Shingeki no Kyojin"
    ],
    "images": {
        "1": "https://cdn.myanimelist.net/images/anime/5/44560l.jpg"
    }
  },
  {
    "id": "625798abc15bd223b8a9eeae10",
    "name": "Attack on Titan",
    "aliases": [
        "Shingeki no Kyojin"
    ],
    "images": {
        "1": "https://cdn.myanimelist.net/images/anime/5/44560l.jpg"
    }
  },
  {
    "id": "625798abc15bd223b8a9eeae11",
    "name": "Attack on Titan",
    "aliases": [
        "Shingeki no Kyojin"
    ],
    "images": {
        "1": "https://cdn.myanimelist.net/images/anime/5/44560l.jpg"
    }
  },
  {
    "id": "625798abc15bd223b8a9eeae12",
    "name": "Attack on Titan",
    "aliases": [
        "Shingeki no Kyojin"
    ],
    "images": {
        "1": "https://cdn.myanimelist.net/images/anime/5/44560l.jpg"
    }
  },
  {
    "id": "625798abc15bd223b8a9eeae13",
    "name": "Attack on Titan",
    "aliases": [
        "Shingeki no Kyojin"
    ],
    "images": {
        "1": "https://cdn.myanimelist.net/images/anime/5/44560l.jpg"
    }
  },
  {
    "id": "625798abc15bd223b8a9eeae14",
    "name": "Attack on Titan",
    "aliases": [
        "Shingeki no Kyojin"
    ],
    "images": {
        "1": "https://cdn.myanimelist.net/images/anime/5/44560l.jpg"
    }
  },
  {
    "id": "625798abc15bd223b8a9eeae15",
    "name": "Attack on Titan",
    "aliases": [
        "Shingeki no Kyojin"
    ],
    "images": {
        "1": "https://cdn.myanimelist.net/images/anime/5/44560l.jpg"
    }
  },
  {
    "id": "625798abc15bd223b8a9eeae16",
    "name": "Attack on Titan",
    "aliases": [
        "Shingeki no Kyojin"
    ],
    "images": {
        "1": "https://cdn.myanimelist.net/images/anime/5/44560l.jpg"
    }
  },
  {
    "id": "625798abc15bd223b8a9eeae17",
    "name": "Attack on Titan",
    "aliases": [
        "Shingeki no Kyojin"
    ],
    "images": {
        "1": "https://cdn.myanimelist.net/images/anime/5/44560l.jpg"
    }
  },
  {
    "id": "625798abc15bd223b8a9eeae18",
    "name": "Attack on Titan",
    "aliases": [
        "Shingeki no Kyojin"
    ],
    "images": {
        "1": "https://cdn.myanimelist.net/images/anime/5/44560l.jpg"
    }
  },
  {
    "id": "625798abc15bd223b8a9eeae19",
    "name": "Attack on Titan",
    "aliases": [
        "Shingeki no Kyojin"
    ],
    "images": {
        "1": "https://cdn.myanimelist.net/images/anime/5/44560l.jpg"
    }
  }
]

function App() {
  return (
    <div className="App">
      <Row>
        <Col>
          <h1>Popit</h1>
          <p>The site for spoiler-free show information</p>
          <br /><br /><br /><br />
        </Col>
      </Row>
      <Row>
        <Col sm={3} />
        <Col sm={6}>
          <Card>
            <CardBody>
              <CardTitle>Find a show!</CardTitle>
              <InputGroup>
                <Input />
                <Button color="primary" id="search"><Spinner color="light" size="sm" />Search!</Button>
              </InputGroup>
            </CardBody>
          </Card>
        </Col>
        <Col sm={3} />
      </Row>
      <UncontrolledCollapse toggler="#search">
      <Row style={{marginTop: '2%'}}>
        <PageScrollMenu name="Results" items={items.map((show, index) => (
          <PageCard itemId={"results_" + show.id} key={"results_" + show.id} src={show.images[1]} title={show.name} 
          nav={"/show/" + show.id} />
        ))} />
      </Row>
      </UncontrolledCollapse>
      <Row style={{marginTop: '2%'}}>
        <PageScrollMenu name="Your Shows" items={items.map((show, index) => (
          <PageCard itemId={"yours_" + show.id} key={"yours_" + show.id} src={show.images[1]} title={show.name}
          nav={"/show/" + show.id} />
        ))} />
      </Row>
      <Row style={{marginTop: '2%'}}>
        <PageScrollMenu name="Browse" items={items.map((show, index) => (
          <PageCard itemId={"browse_" + show.id} key={"browse_" + show.id} src={show.images[1]} title={show.name}
          nav={"/show/" + show.id} />
        ))} />
      </Row>
      <Link to={"/show/" + items[0].id}>LINK</Link>
    </div>
  );
}

export default App;
