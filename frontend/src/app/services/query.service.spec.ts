import { TestBed } from '@angular/core/testing';
import { QueryParser } from './query.service'; // Adjust the path as needed
import { equal, contains, startsWith, field } from './generic.service'; // Adjust the path as needed

describe('QueryParser', () => {
	let parser: QueryParser;

	beforeEach(() => {
		TestBed.configureTestingModule({});
		parser = new QueryParser();
	});

	it('should parse equal conditions correctly', () => {
		const queryString = 'name:John age:30';
		const query = parser.parse(queryString);

		expect(query.conditions).toEqual([
			equal(field('name'), 'John'),
			equal(field('age'), 30),
		]);
	});

	it('should parse contains conditions correctly', () => {
		const queryString = 'description:~keyword';
		const query = parser.parse(queryString);

		expect(query.conditions).toEqual([
			contains(field('description'), 'keyword'),
		]);
	});

	it('should parse starts with conditions correctly', () => {
		const queryString = 'name:^John';
		const query = parser.parse(queryString);

		expect(query.conditions).toEqual([startsWith(field('name'), 'John')]);
	});

	it('should parse mixed conditions correctly', () => {
		const queryString = 'name:^John description:~keyword age:30';
		const query = parser.parse(queryString);

		expect(query.conditions).toEqual([
			startsWith(field('name'), 'John'),
			contains(field('description'), 'keyword'),
			equal(field('age'), 30),
		]);
	});

	it('should parse order by correctly', () => {
		const queryString = 'orderBy:name:asc,age:desc';
		const query = parser.parse(queryString);

		expect(query.orderBy).toEqual([
			{ field: 'name', desc: false },
			{ field: 'age', desc: true },
		]);
	});

	it('should parse limit correctly', () => {
		const queryString = 'limit:10';
		const query = parser.parse(queryString);

		expect(query.limit).toBe(10);
	});

	it('should parse after correctly', () => {
		const queryString = 'after:id1,id2';
		const query = parser.parse(queryString);

		expect(query.after).toEqual(['id1', 'id2']);
	});

	it('should parse a full query string correctly', () => {
		const queryString =
			'name:^John description:~keyword age:30 orderBy:name:asc,age:desc limit:10 after:id1,id2';
		const query = parser.parse(queryString);

		expect(query.conditions).toEqual([
			startsWith(field('name'), 'John'),
			contains(field('description'), 'keyword'),
			equal(field('age'), 30),
		]);
		expect(query.orderBy).toEqual([
			{ field: 'name', desc: false },
			{ field: 'age', desc: true },
		]);
		expect(query.limit).toBe(10);
		expect(query.after).toEqual(['id1', 'id2']);
	});

	it('should parse default query correctly', () => {
		const queryString = 'ASD';
		const query = parser.parse(queryString);

		expect(query.conditions).toEqual([contains(field('name'), 'ASD')]);
	});

	it('should parse multifield query correctly', () => {
		const queryString = 'email,name:~ASD';
		const query = parser.parse(queryString);

		expect(query.conditions).toEqual([
			contains(field('email'), 'ASD'),
			contains(field('name'), 'ASD'),
		]);
	});
});
