import React, { useState } from 'react';
import { Card, CardImg, CardImgOverlay, CardTitle } from "reactstrap"
import { useNavigate } from "react-router-dom";


export function PageCard(props) {
    const [onHover, setOnHover] = useState(false);
    const navigate = useNavigate();

    return (
      <Card style={{display: 'inline-block', margin: '5px 5px', width: '160px', height: '220px', cursor: 'pointer'}} 
      onMouseEnter={() => setOnHover(true)} onMouseLeave={() => setOnHover(false)}
      onClick={() => navigate(props.nav)}>
        <CardImg top src={props.src} width="100%" />
        <CardImgOverlay style={{backgroundColor: (onHover ? 'rgba(0, 0, 0, 0.9)' : 'rgba(0, 0, 0, 0)')}}>
          <CardTitle style={{ color: "white" }}>{onHover ? props.title : ""}</CardTitle>
        </CardImgOverlay>
      </Card>
    );
}
