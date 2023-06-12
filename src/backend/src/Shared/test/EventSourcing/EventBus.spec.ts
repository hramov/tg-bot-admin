import { Event } from 'src/Shared/src/EventSourcing/Event';
import { EventBus } from 'src/Shared/src/EventSourcing/EventBus';
import { ICommand } from 'src/Shared/src/EventSourcing/ICommand';
import { Uuid } from 'src/Shared/src/ValueObject/Objects/Uuid';

describe('EventBus', () => {
	it('Test general behavior', () => {
		class TestCommand implements ICommand {
			public id: Uuid;
			public data: string;
			constructor(data: string) {
				this.id = new Uuid();
			}
		}

		const bus = new EventBus();

		bus.listenTo(
			new Event<TestCommand>((data: any) => {
				expect(data).toBe('test');
			}),
		);

		const command = new TestCommand('test');
		bus.sendCommand(command);
	});
});
