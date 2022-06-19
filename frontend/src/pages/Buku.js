import React, {useState, Link} from 'react';
import { Button, Card, Col, Container, Row } from "react-bootstrap"


function BasicExample(props) {
  return (
    <Container>
      <Row xs={4} className="g-4">
        {Array.from({ length: 8 }).map((_, idx) => (
          <Col>
            <Card style={{ width: '15rem'}}>
              <Card.Img variant="top" src={props.imgSrc} />
              <Card.Body>
                <Card.Title>{props.title}</Card.Title>
                <Card.Text>
                  {props.text}
                </Card.Text>
                <Button variant="primary">Go somewhere</Button>
              </Card.Body>
            </Card>
          </Col>
        ))}
        </Row>
    </Container>
  );
}

export default BasicExample;