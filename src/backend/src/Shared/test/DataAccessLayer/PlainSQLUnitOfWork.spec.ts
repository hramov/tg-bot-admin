import { PostgresTransactionStorage } from 'src/App/Infrastructure/PersistentStorage/Storage/Postgres/PostgresTransaction.storage';
import { IDatabaseConnection } from 'src/Shared/src/DataAccessLayer/IDatabaseConnection';
import { PlainSQLUnitOfWork } from 'src/Shared/src/DataAccessLayer/UnitOfWork/PlainSQLUnitOfWork';

describe('PlanSQLUnitOfWork', () => {
	describe('query', () => {
		it('', () => {});
		// it('should throw connection undefined error', async () => {
		// 	const connection: IDatabaseConnection = null;
		// 	const unitOfWork = new PlainSQLUnitOfWork(connection);

		// 	expect(await unitOfWork.query(['select 1'])).toThrowError(
		// 		'Connection is undefined',
		// 	);
		// });

		// it('should return select 1', async () => {
		// 	const connection: IDatabaseConnection =
		// 		new PostgresTransactionStorage({
		// 			host: 'localhost',
		// 			port: 5432,
		// 			client: 'postgres',
		// 			password: 'postgres',
		// 			database: 'test',
		// 		});
		// 	const unitOfWork = new PlainSQLUnitOfWork(connection);

		// 	const result = await unitOfWork.query(['select 1']);
		// 	console.log(result);
		// });
	});
});
