import { Module } from '@nestjs/common';
import { AppController } from '@/infra/nest/controllers';
import { AppService } from '@/infra/nest/services';

@Module({
  imports: [],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
