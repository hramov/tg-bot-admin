import {Admin} from "../../src/Core/Context/Admin/Admin";
import {IEventPublisher} from "../../src/Shared/src/EventSourcing/IEventPublisher";
import {Status} from "../../src/Core/Context/Admin/Entity/Status";
import {TnkType} from "../../src/Core/Context/Admin/Entity/TnkType";

describe('Admin aggregate test', () => {
    describe('calculateDelta', () => {
        it('should return one field that is different', () => {
            const eventBus: IEventPublisher = null;
            const tnk = new Admin(eventBus);

            tnk.load({
                title: 'Название',
                isActive: true,
                isDigital: true,
                isAutomated: true,
                status: new Status('', 30),
                type: new TnkType('общий'),
                processId: 34,
                subprocessId: 56,
            })

            const diff = tnk.calculateDelta({
                title: 'Название 1',
                isActive: true,
                isDigital: true,
                isAutomated: true,
                status: new Status('', 30),
                type: new TnkType('общий'),
                processId: 34,
                subprocessId: 56,
            });

            expect(diff).toStrictEqual({title: 'Название 1'});
        })
    })
})