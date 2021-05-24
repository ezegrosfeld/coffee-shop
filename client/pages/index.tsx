import Head from 'next/head';
import useSWR from 'swr';
import styles from '../styles/Home.module.css';

export default function Home() {
	const { data: products, error } = useSWR(
		'http://172.24.101.211:8080/products'
	);
	console.log(products);
	return (
		<div className={styles.container}>
			<Head>
				<title>Create Next App</title>
				<meta name="description" content="Generated by create next app" />
				<link rel="icon" href="/favicon.ico" />
			</Head>

			<main className={styles.main}>
				<ul>
					{products?.map((product) => (
						<li key={product.id}>{product.name}</li>
					))}
				</ul>
			</main>
		</div>
	);
}