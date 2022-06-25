import React, {useState, Link} from 'react';
import ModalFooter from 'react-bootstrap/ModalFooter'
import {
Row,
Column,
Box,
Container,
FooterLink,
Heading,
} from "./FooterStyles";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'

const Footer = () => {
return (
	<Box>
		<Row xs={4} className="">
		<Column>
			<FooterLink href="#">
			<FontAwesomeIcon icon="fa-brands fa-facebook-f" />
				<span style={{ marginLeft: "5px" }}>
				Budaya Baca
				</span>
			</FooterLink>
		</Column>
		<Column>
			<FooterLink href="#">
			<FontAwesomeIcon icon="coffee" />
			<i className="fab fa-instagram">
				<span style={{ marginLeft: "5px" }}>
				@budaya_baca
				</span>
			</i>
			</FooterLink>
		</Column>
		<Column>
			<FooterLink href="#">
			<i className="fab fa-twitter">
				<span style={{ marginLeft: "5px" }}>
				budaya_baca
				</span>
			</i>
			</FooterLink>
		</Column>
		<Column>
			<FooterLink href="#">
			<FontAwesomeIcon icon={['fab', 'apple']} />
			<i className="fab fa-youtube">
				<span style={{ marginLeft: "5px" }}>
				Budaya Baca
				</span>
			</i>
			</FooterLink>
		</Column>
		</Row>
    </Box>
);
};
export default Footer;
