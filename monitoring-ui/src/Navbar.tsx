//import React, { useEffect, useState } from 'react'
import styled from 'styled-components'

import { Inter700 } from './components/Styles/Styles'

function Navbar() {

	return (
		<NavbarSection>
			<Container>
				<div aria-hidden='true'>
					<img src='/assets/dappnode_logo.png' alt='logo' height={50}/>
					<Inter700>Pinner DAO</Inter700>
				</div>
			</Container>
		</NavbarSection>
	)
}

// eslint-disable-next-line no-unused-vars
/*const ConnectedLine = styled.div`
	background: #54d4cb;
	box-shadow: 0px 0px 12px #54d4cb;
	width: 100%;
	height: 4px;
`
*/

const NavbarSection = styled.section`
	font-family: Roboto, Helvetica, sans-serif;
	background-color: transparent;
	height: 72px;
	@media only screen and (max-width: 768px) {
		height: auto;
	}
	text-align: center;
	overflow: hidden;
	position: relative;
	z-index: 1;
`

const Container = styled.div`
	width: 95%;
	margin: auto;
	height: 100%;
	display: flex;
	align-items: center;
	justify-content: space-between;
	flex-direction: row;
	@media only screen and (max-width: 768px) {
		flex-direction: column;
	}
	h1 {
		font-family: Roboto, Helvetica, sans-serif;
		font-size: 24px;
		display: flex;
		align-items: center;
		letter-spacing: 0.2px;
		margin-left: 20px;
		img {
			padding-right: 32px;
		}
	}
	div {
		display: flex;
		align-items: center;
		justify-content: space-between;
		flex-direction: row;
		cursor: pointer;
		p {
			margin: 0 15px;
		}
	}
`

export default Navbar
