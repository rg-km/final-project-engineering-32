import React, {useState, Link} from 'react';
import { Col,Row,Form,Button, Container } from 'react-bootstrap'

function Pinjam_buku() {

    return (
  
<Form>
<Container>
  <Row className="mb-3">
    <Form.Group as={Col} controlId="formGridEmail">
      <Form.Label>Tgl Peminjaman</Form.Label>
      <Form.Control type="Tgl Peminjaman" placeholder="Enter Date" />
    </Form.Group>

    <Form.Group as={Col} controlId="formGridPassword">
      <Form.Label>Tgl Kembali</Form.Label>
      <Form.Control type="Tgl Kembali" placeholder="Enter Date" />
    </Form.Group>
  </Row>

  <Form.Group className="mb-3" controlId="formGridAddress1">
    <Form.Label>Nama</Form.Label>
    <Form.Control placeholder="Nama" />
  </Form.Group>

  <Form.Group className="mb-3" controlId="formGridAddress2">
    <Form.Label>No Transaksi</Form.Label>
    <Form.Control placeholder="No Transaksi" />
  </Form.Group>

  <Row className="mb-3">
    <Form.Group as={Col} controlId="formGridCity">
      <Form.Label>Kode Buku</Form.Label>
      <Form.Control />
    </Form.Group>


    <Form.Group as={Col} controlId="formGridZip">
      <Form.Label>Judul Buku</Form.Label>
      <Form.Control />
    </Form.Group>
  </Row>


  <Button variant="primary" type="simpan">
    Submit
  </Button>
  </Container>
</Form>

    )
    
}
export default Pinjam_buku;
